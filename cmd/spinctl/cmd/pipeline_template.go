// Copyright 2019 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

type pipelineTemplate struct {
	ioStreams  *genericclioptions.IOStreams
	client     *spinnaker.Client
	output     string
	listScopes []string
}

// NewCmdPipelineTemplate manage the pipeline template command.
func NewCmdPipelineTemplate(ctx context.Context, client *spinnaker.Client, ioStreams *genericclioptions.IOStreams) *cobra.Command {
	pt := &pipelineTemplate{
		ioStreams: ioStreams,
		client:    client,
	}

	cmd := &cobra.Command{
		Use:   "pipeline-template",
		Short: "manage the Spinnaker pipeline template.",
	}
	cmd.AddCommand(pt.get(ctx))
	cmd.AddCommand(pt.list(ctx))
	cmd.AddCommand(pt.create(ctx))
	cmd.AddCommand(pt.update(ctx))
	cmd.AddCommand(pt.delete(ctx))

	return cmd
}

func (pt *pipelineTemplate) get(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the specified pipeline template.",
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 1); err != nil {
				return err
			}

			ctx, err = pt.client.Authenticate(ctx)
			return err
		},
		RunE: func(_ *cobra.Command, args []string) error {
			id := args[0]
			logger.FromContext(ctx).Debugf("pipelineTemplate.get: id: %s", id)

			s, err := pt.client.GetPipelineTemplate(ctx, id, pt.output)
			if err != nil {
				return err
			}
			fmt.Fprintf(pt.ioStreams.Out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&pt.output, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

func (pt *pipelineTemplate) list(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List pipeline templates.",
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 0); err != nil {
				return err
			}

			ctx, err = pt.client.Authenticate(ctx)
			return err
		},
		RunE: func(*cobra.Command, []string) error {
			s, err := pt.client.ListPipelineTemplates(ctx, pt.listScopes, pt.output)
			if err != nil {
				return err
			}
			fmt.Fprintf(pt.ioStreams.Out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringArrayVar(&pt.listScopes, "scopes", nil, "pipeline template scopes")
	f.StringVarP(&pt.output, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

func (pt *pipelineTemplate) create(context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create a pipeline template.",
		PreRunE: func(*cobra.Command, []string) (err error) { return nil },
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}

func (pt *pipelineTemplate) update(context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"up"},
		Short:   "Update a pipeline template.",
		Args:    cobra.ExactArgs(1),
		PreRunE: func(*cobra.Command, []string) (err error) { return nil },
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}

func (pt *pipelineTemplate) delete(context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete a pipeline template.",
		PreRunE: func(*cobra.Command, []string) (err error) { return nil },
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}
