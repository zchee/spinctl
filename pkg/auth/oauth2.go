// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"bufio"
	"context"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	sha256 "github.com/minio/sha256-simd"
	"github.com/pkg/errors"
	open "github.com/zchee/go-open"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	oauth2_google "golang.org/x/oauth2/google"
	oauth2_v2 "google.golang.org/api/oauth2/v2"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/unsafeutil"
)

// OAuth2Config implements a OAuth2.0 authentication data for Spinnaker.
type OAuth2Config struct {
	ClientID     string        `yaml:"clientId,omitempty"`
	ClientSecret string        `yaml:"clientSecret,omitempty"`
	AuthURL      string        `yaml:"authUrl,omitempty"`
	TokenURL     string        `yaml:"tokenUrl,omitempty"`
	Scopes       []string      `yaml:"scopes,omitempty"`
	Token        *oauth2.Token `yaml:"token,omitempty"`
}

// IsValid checks whethere the OAuth2Config is valid.
func (o OAuth2Config) IsValid() bool {
	return o.ClientID != "" && o.ClientSecret != "" && o.TokenURL != "" && o.AuthURL != "" && len(o.Scopes) != 0
}

// AuthenticateOAuth2 authenticate gate with OAuth2 and returns the new oauth2.Token.
//
// Currently, only of supported Google OAuth2 service.
//  https://developers.google.com/identity/protocols/OpenIDConnect
func AuthenticateOAuth2(ctx context.Context, oc *OAuth2Config) (*oauth2.Token, error) {
	const oauth2CallbackAddr = ":8085" // TODO(zchee): avoid set localhost to Google OAuth2 Authorized redirect URIs

	conf := &oauth2.Config{
		ClientID:     oc.ClientID,
		ClientSecret: oc.ClientSecret,
		RedirectURL:  "http://localhost" + oauth2CallbackAddr,
		Scopes:       oc.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  oc.AuthURL,
			TokenURL: oc.TokenURL,
		},
	}

	// TODO(zchee): Currently, only of supported Google OAuth2 if Scope is nil
	if conf.Scopes == nil {
		conf.Scopes = []string{"openid", oauth2_v2.UserinfoEmailScope, oauth2_v2.UserinfoProfileScope}
	}
	if conf.Endpoint.AuthURL == "" || conf.Endpoint.TokenURL == "" {
		conf.Endpoint = oauth2_google.Endpoint
	}

	// Requests refresh token
	authCodeOps := []oauth2.AuthCodeOption{oauth2.AccessTypeOffline}

	if strings.EqualFold(conf.Endpoint.AuthURL, "https://accounts.google.com/o/oauth2/auth") {
		authCodeOps = append(authCodeOps,
			oauth2.ApprovalForce, // Use `approval_prompt=force` for Google OAuth2
			oauth2.SetAuthURLParam("include_granted_scopes", "true"), // `include_granted_scopes` values is google specific
		)
	} else {
		// The `oauth2.ApprovalForce` is Google specific. It's old value and not RFC.
		// Use `prompt=select_account` instead of `approval_prompt=force` if *not* Google OAuth2 service.
		authCodeOps = append(authCodeOps, oauth2.SetAuthURLParam("prompt", "select_account"))
	}

	// Support RFC7636, Authorization code flow with PKCE
	//  https://tools.ietf.org/html/rfc7636
	state, codeChallenge, err := StateAndCodeChallenge()
	if err != nil {
		return nil, errors.Wrap(err, "could not create state and codeChallenge strings")
	}

	const challengeMethodS256 = "S256"
	authCodeOps = append(authCodeOps,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", challengeMethodS256),
	)
	authURL := conf.AuthCodeURL(state, authCodeOps...)

	cs := CallbackServer(oauth2CallbackAddr)
	go cs.ListenAndServe()
	defer cs.Close()

	live := make(chan struct{}, 1)
	go func() {
		var tmpDelay time.Duration // how long to sleep on accept failure
		for {
			if _, err = net.Dial("tcp", oauth2CallbackAddr); err != nil {
				logger.FromContext(ctx).Error("Dial error", zap.Duration("retrying", tmpDelay), zap.Error(err))

				if tmpDelay == 0 {
					tmpDelay = 5 * time.Millisecond
				} else {
					tmpDelay *= 2
				}
				if max := 1 * time.Second; tmpDelay > max {
					tmpDelay = max
				}
				timer := time.NewTimer(tmpDelay)
				select {
				case <-timer.C:
				}
				continue
			}
			break
		}

		live <- struct{}{}
		logger.FromContext(ctx).Debug("succsess Dialing")
	}()
	<-live // wait for callbackServer is live

	// Opens Google account login page
	fmt.Fprintf(os.Stdout, "Your browser has been opened to visit:\n\n\t%s\n", authURL)
	if err := open.RunContext(ctx, authURL); err != nil {
		return nil, err
	}

	code := CodeFromStdin()
	if code == "" {
		return nil, errors.New("auth: authorization code is empty")
	}

	tok, err := conf.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", state)) // set code_verifier param
	if err != nil {
		return nil, errors.Wrap(err, "auth: failed to converts an authorization code into a token")
	}

	return tok, nil
}

// Login login to gate.
func Login(hc *http.Client, oc *OAuth2Config, gateEndpoint string, token *oauth2.Token) error {
	req, err := http.NewRequest(http.MethodGet, gateEndpoint+"/login", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	if _, err := hc.Do(req); err != nil {
		return err
	}

	return nil
}

// CallbackServer returns the new http.Server which dummy http server for OAuth2 callback handler.
func CallbackServer(addr string) *http.Server {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.FormValue("code"))
	})

	return &http.Server{
		Addr:    addr,
		Handler: h,
	}
}

// CodeFromStdin waits to users stdin input and returns the input with trim the spaces.
func CodeFromStdin() string {
	r := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Paste authorization code: ")
	code, err := r.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(code)
}

// StateAndCodeChallenge implements the OAth PKCE RFC7636 specification.
//  https://tools.ietf.org/html/rfc7636#section-4.2
//
// The returns state(code_verifier) and code_challenge.
// state is base64 encoded random string.
// codeChallenge is sha256 hashed and base64 encoded state string.
func StateAndCodeChallenge() (state, codeChallenge string, err error) {
	b := make([]byte, 1024)
	if _, err := crand.Read(b); err != nil {
		return "", "", errors.Wrap(err, "could not generate random string")
	}

	state = base64.RawURLEncoding.EncodeToString(b)
	hv := HashVerifier(unsafeutil.UnsafeSlice(state))
	codeChallenge = base64.RawURLEncoding.EncodeToString(hv)

	return state, codeChallenge, nil
}

// HashVerifier creates b SHA256 checksum and returns sum in bytes.
//
// If host CPU supporting Simd instruction, Use Ssse, Avx, Avx2 or Avx512.
func HashVerifier(b []byte) []byte {
	h := sha256.New()
	h.Write(b)

	return h.Sum(nil)
}
