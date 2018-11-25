// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"io"

	"github.com/spf13/cobra"

	"github.com/zchee/spinctl/pkg/spinnaker"
)

func NewCmdVersion(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	get := &applicationGet{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get Gate's current version",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateArgs(args); err != nil {
				return err
			}

			return client.GetVersion(ctx, get.out, get.outFormat)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&get.outFormat, "output", "o", "", "output format. One of: (json|yaml)")

	return cmd
}
