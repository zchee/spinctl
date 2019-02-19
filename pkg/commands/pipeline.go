// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commands

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/zchee/spinctl/pkg/logging"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

type pipeline struct {
	ioStreams    *genericclioptions.IOStreams
	client       *spinnaker.Client
	outputFormat string
}

// NewCmdPipeline creates the pipeline command.
func NewCmdPipeline(ctx context.Context, client *spinnaker.Client, ioStreams *genericclioptions.IOStreams) *cobra.Command {
	p := &pipeline{
		ioStreams: ioStreams,
		client:    client,
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
		Use:   "get [application name] [pipeline name]",
		Short: "Get the specified pipeline.",
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 2); err != nil {
				return err
			}

			ctx, err = p.client.Authenticate(ctx)
			return err
		},
		RunE: func(_ *cobra.Command, args []string) error {
			application := args[0]
			pipelineName := args[1]
			logging.FromContext(ctx).Debugf("CmdPipelineGet: application: %s, pipelineName: %s", application, pipelineName)

			s, err := p.client.GetPipeline(ctx, application, pipelineName, p.outputFormat)
			if err != nil {
				return err
			}
			fmt.Fprintf(p.ioStreams.Out, s)
			return nil
		},
	}

	addPrintFlags(&p.outputFormat).AddFlags(cmd)

	return cmd
}

func (p *pipeline) list(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list <application name>",
		Aliases: []string{"ls"},
		Short:   "List the pipelines for the provided application.",
		Example: "  spinctl pipeline list spin -o yaml",
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 2); err != nil {
				return err
			}
			ctx, err = p.client.Authenticate(ctx)
			return err
		},
		RunE: func(_ *cobra.Command, args []string) error {
			name := args[0]
			logging.FromContext(ctx).Debugf("CmdPipelineList: name: %s", name)

			s, err := p.client.ListPipelines(ctx, name, p.outputFormat)
			if err != nil {
				return err
			}
			fmt.Fprintf(p.ioStreams.Out, s)
			return nil
		},
	}

	addPrintFlags(&p.outputFormat).AddFlags(cmd)

	return cmd
}

func (p *pipeline) save(context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "save",
		Short:   "Save the provided pipeline.",
		PreRunE: func(*cobra.Command, []string) (err error) { return nil },
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}

func (p *pipeline) delete(context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete the provided pipeline.",
		PreRunE: func(*cobra.Command, []string) (err error) { return nil },
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}

func (p *pipeline) execute(context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "execute",
		Aliases: []string{"exec"},
		Short:   "Execute the provided pipeline.",
		PreRunE: func(*cobra.Command, []string) (err error) { return nil },
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}
