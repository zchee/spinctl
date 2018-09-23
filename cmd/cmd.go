// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/zchee/spinctl/api/gate"
	"github.com/zchee/spinctl/pkg/auth"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

func NewDefaultCommand(ctx context.Context, args []string) *cobra.Command {
	return NewCommand(ctx, os.Stdin, os.Stdout, os.Stderr, args)
}

func NewCommand(ctx context.Context, in io.Reader, out, er io.Writer, args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "spinctl",
		Short:        "spinctl is a command-line tool to manage Spinnaker via gate API.",
		SilenceUsage: true,
	}

	flags := cmd.PersistentFlags()
	flags.Parse(args)

	tok, err := auth.AuthenticateOAuth2(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !tok.Valid() {
		log.Fatalf("token is invalid: %v", tok)
	}
	ctx = context.WithValue(ctx, gate.ContextAccessToken, tok.AccessToken)

	client := spinnaker.NewClient()
	cmd.AddCommand(NewCmdApplication(ctx, client))

	return cmd
}
