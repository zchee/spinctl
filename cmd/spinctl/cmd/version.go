// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/zchee/spinctl/pkg/spinnaker"
)

type version struct {
	out    io.Writer
	client *spinnaker.Client
	output string
}

// NewCmdVersion creates the version command.
func NewCmdVersion(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	v := &version{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get Gate's current version",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = v.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateArgs(args); err != nil {
				return err
			}

			s, err := client.GetVersion(ctx, v.output)
			if err != nil {
				return err
			}
			fmt.Fprintf(v.out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&v.output, "output", "o", "", "output format. One of: (json|yaml)")

	return cmd
}
