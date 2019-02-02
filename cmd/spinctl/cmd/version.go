// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/zchee/spinctl/pkg/spinnaker"
)

type version struct {
	ioStreams *genericclioptions.IOStreams
	client    *spinnaker.Client
	output    string
}

// NewCmdVersion creates the version command.
func NewCmdVersion(ctx context.Context, client *spinnaker.Client, ioStreams *genericclioptions.IOStreams) *cobra.Command {
	v := &version{
		ioStreams: ioStreams,
		client:    client,
	}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show current Gate's version",
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
			fmt.Fprintf(v.ioStreams.Out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&v.output, "output", "o", "", "output format. One of: (json|yaml)")

	return cmd
}
