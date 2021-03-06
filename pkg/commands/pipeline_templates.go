// Copyright 2019 The spinctl Authors. All rights reserved.
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

type pipelineTemplates struct {
	ioStreams    *genericclioptions.IOStreams
	client       *spinnaker.Client
	outputFormat string

	getVersion string
	getDigest  string

	listScopes []string

	saveFile string
}

// NewCmdPipelineTemplates manage the pipeline template command.
func NewCmdPipelineTemplates(ctx context.Context, client *spinnaker.Client, ioStreams *genericclioptions.IOStreams) *cobra.Command {
	pt := &pipelineTemplates{
		ioStreams: ioStreams,
		client:    client,
	}

	cmd := &cobra.Command{
		Use:   "pipeline-templates",
		Short: "manage the Spinnaker pipeline templates.",
	}
	cmd.AddCommand(pt.get(ctx))
	cmd.AddCommand(pt.list(ctx))
	cmd.AddCommand(pt.save(ctx))
	cmd.AddCommand(pt.update(ctx))
	cmd.AddCommand(pt.delete(ctx))

	return cmd
}

func (pt *pipelineTemplates) get(ctx context.Context) *cobra.Command {
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
			logging.FromContext(ctx).Debugf("pipelineTemplates.get: id: %s", id)

			opts := spinnaker.GetPipelineTemplateOptions{
				Version: pt.getVersion,
				Digest:  pt.getDigest,
			}
			s, err := pt.client.GetPipelineTemplate(ctx, pt.outputFormat, id, opts)
			if err != nil {
				return err
			}
			fmt.Fprintf(pt.ioStreams.Out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&pt.getVersion, "template-version", "V", "", "Version to pipeline-templates")
	f.StringVarP(&pt.getDigest, "template-Digest", "D", "", "Digest to pipeline-templates")
	addPrintFlags(&pt.outputFormat).AddFlags(cmd)

	return cmd
}

func (pt *pipelineTemplates) list(ctx context.Context) *cobra.Command {
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
			s, err := pt.client.ListPipelineTemplates(ctx, pt.listScopes, pt.outputFormat)
			if err != nil {
				return err
			}
			fmt.Fprintf(pt.ioStreams.Out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringArrayVar(&pt.listScopes, "scopes", nil, "pipeline template scopes")
	addPrintFlags(&pt.outputFormat).AddFlags(cmd)

	return cmd
}

func (pt *pipelineTemplates) save(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save a pipeline template.",
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 0); err != nil {
				return err
			}

			ctx, err = pt.client.Authenticate(ctx)
			return err
		},
		RunE: func(*cobra.Command, []string) error {
			if err := pt.client.SavePipelineTemplates(ctx, pt.saveFile); err != nil {
				return err
			}
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&pt.saveFile, "file", "f", "", "file path to pipeline-template json")

	return cmd
}

func (pt *pipelineTemplates) update(context.Context) *cobra.Command {
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

func (pt *pipelineTemplates) delete(context.Context) *cobra.Command {
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
