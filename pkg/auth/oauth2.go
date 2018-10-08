// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	oauth2_google "golang.org/x/oauth2/google"
	oauth2_v2 "google.golang.org/api/oauth2/v2"
)

// OAuth2Config implements a OAuth2.0 authentication data for Spinnaker.
type OAuth2Config struct {
	ClientID     string        `yaml:"clientId"`
	ClientSecret string        `yaml:"clientSecret"`
	AuthURL      string        `yaml:"authUrl"`
	TokenURL     string        `yaml:"tokenUrl"`
	Scopes       []string      `yaml:"scopes"`
	Token        *oauth2.Token `yaml:"token,omitempty"`
}

// IsValid checks whethere the OAuth2Config is valid.
func (o OAuth2Config) IsValid() bool {
	return o.ClientID != "" && o.ClientSecret != "" && o.TokenURL != "" && o.AuthURL != "" && len(o.Scopes) != 0
}

func AuthenticateOAuth2(ctx context.Context) (*oauth2.Token, error) {
	const oauth2CallbackAddr = ":8085"
	conf := &oauth2.Config{
		ClientID:     os.Getenv("SPINCTL_OAUTH2_CLIENT_ID"),
		ClientSecret: os.Getenv("SPINCTL_OAUTH2_CLIENT_SECRET"),
		RedirectURL:  "http://localhost" + oauth2CallbackAddr,                                // TODO(zchee): avoid set localhost to Google OAuth2 Authorized redirect URIs
		Scopes:       []string{oauth2_v2.UserinfoEmailScope, oauth2_v2.UserinfoProfileScope}, // TODO(zchee): currently only supported Google OAuth2
		Endpoint:     oauth2_google.Endpoint,
	}

	// for OAuth2 callback handler
	cs := &http.Server{
		Addr: oauth2CallbackAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO(zchee): state validation
			fmt.Fprintln(w, r.FormValue("code"))
		}),
	}
	go cs.ListenAndServe()
	defer cs.Close()

	// TODO(zchee): support RFC 7636, authorization code flow with PKCE
	//  https://tools.ietf.org/html/rfc7636
	//  https://www.oauth.com/oauth2-servers/pkce/authorization-code-exchange/
	//  https://developer.okta.com/authentication-guide/implementing-authentication/auth-code-pkce
	authURL := conf.AuthCodeURL("state",
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
		oauth2.SetAuthURLParam("include_granted_scopes", "true"), // TODO(zchee): include_granted_scopes is google specific?
		// TODO(zchee): set code_challenge param
	)
	fmt.Fprintf(os.Stdout, "Your browser has been opened to visit:\n\n\t%s\n", authURL)
	code := codeFromStdin()
	if code == "" {
		return nil, errors.New("auth: authorization code is empty")
	}

	tok, err := conf.Exchange(ctx, code) // TODO(zchee): set code_verifier param
	if err != nil {
		return nil, errors.Wrap(err, "auth: failed to converts an authorization code into a token")
	}

	return tok, nil
}

func codeFromStdin() string {
	r := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Paste authorization code: ")
	code, _ := r.ReadString('\n')

	return strings.TrimSpace(code)
}
