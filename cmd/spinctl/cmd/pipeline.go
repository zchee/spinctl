// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

type pipeline struct {
	out    io.Writer
	client *spinnaker.Client
	output string
}

// NewCmdPipeline creates the pipeline command.
func NewCmdPipeline(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	p := &pipeline{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   "pipeline",
		Short: "manage the Spinnaker pipelines.",
	}
	cmd.AddCommand(p.delete(ctx))
	cmd.AddCommand(p.execute(ctx))
	cmd.AddCommand(p.get(ctx))
	cmd.AddCommand(p.list(ctx))
	cmd.AddCommand(p.save(ctx))

	return cmd
}

func (p *pipeline) get(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the specified pipeline.",
		Args:  cobra.ExactArgs(2),
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = p.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateArgs(args); err != nil {
				return err
			}

			application := args[0]
			pipelineName := args[1]
			logger.FromContext(ctx).Debugf("CmdPipelineGet: application: %s, pipelineName: %s", application, pipelineName)

			s, err := p.client.GetPipeline(ctx, application, pipelineName, p.output)
			if err != nil {
				return err
			}
			fmt.Fprintf(p.out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&p.output, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

func (p *pipeline) list(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list <application name>",
		Aliases: []string{"ls"},
		Short:   "List the pipelines for the provided application.",
		Args:    cobra.ExactArgs(1),
		Example: "  spinctl pipeline list spin -o yaml",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = p.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateArgs(args); err != nil {
				return err
			}

			name := args[0]
			logger.FromContext(ctx).Debugf("CmdPipelineList: name: %s", name)

			s, err := p.client.ListPipelines(ctx, name, p.output)
			if err != nil {
				return err
			}
			fmt.Fprintf(p.out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&p.output, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

func (p *pipeline) save(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save the provided pipeline.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = p.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}

func (p *pipeline) delete(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete the provided pipeline.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = p.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}

func (p *pipeline) execute(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "execute",
		Aliases: []string{"exec"},
		Short:   "Execute the provided pipeline.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = p.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}
