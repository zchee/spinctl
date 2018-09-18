// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	oauth2_google "golang.org/x/oauth2/google"
	oauth2_v2 "google.golang.org/api/oauth2/v2"

	"github.com/zchee/spinctl/api/gate"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

func NewDefaultCilCommand(args []string) *cobra.Command {
	return NewCommand(os.Stdin, os.Stdout, os.Stderr, args)
}

func NewCommand(sin io.Reader, sout, serr io.Writer, args []string) *cobra.Command {
	ctx := context.Background()

	cmd := &cobra.Command{
		Use:          "spinctl",
		Short:        "spinctl is a command-line tool that management of Spinnaker via gate API.",
		SilenceUsage: true,
	}

	flags := cmd.PersistentFlags()
	flags.Parse(args)

	tok, err := authenticateOAuth2()
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, gate.ContextAccessToken, tok.AccessToken)

	client := spinnaker.New()
	cmd.AddCommand(NewCmdApplication(ctx, client))

	return cmd
}

func authenticateOAuth2() (*oauth2.Token, error) {
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
	code := prompt()

	tok, err := conf.Exchange(context.TODO(), code) // TODO(zchee): set code_verifier param
	if err != nil {
		return nil, errors.Wrap(err, "failed to converts an authorization code into a token")
	}

	return tok, nil
}

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Paste authorization code: ")
	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}
