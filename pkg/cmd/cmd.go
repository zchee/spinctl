// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/zchee/spinctl/pkg/spinnaker"
)

func NewDefaultCommand(ctx context.Context, args []string) *cobra.Command {
	return NewCommand(ctx, args)
}

func NewCommand(ctx context.Context, args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "spinctl",
		Short:        "spinctl is a command-line tool to manage Spinnaker via gate API.",
		SilenceUsage: true,
	}

	flags := cmd.PersistentFlags()
	flags.Parse(args)

	client := spinnaker.NewClient()
	out := cmd.OutOrStdout()

	cmd.AddCommand(NewCmdApplication(ctx, client, out))

	return cmd
}
