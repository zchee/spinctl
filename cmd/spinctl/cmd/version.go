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
	ioStreams    *genericclioptions.IOStreams
	client       *spinnaker.Client
	outputFormat string
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
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 0); err != nil {
				return err
			}

			ctx, err = v.client.Authenticate(ctx)
			return err
		},
		RunE: func(*cobra.Command, []string) error {
			s, err := client.GetVersion(ctx, v.outputFormat)
			if err != nil {
				return err
			}
			fmt.Fprintf(v.ioStreams.Out, s)
			return nil
		},
	}

	addPrintFlags(&v.outputFormat).AddFlags(cmd)

	return cmd
}
