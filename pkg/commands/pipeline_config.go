// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commands

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

type pipelineConfig struct {
	ioStreams    *genericclioptions.IOStreams
	client       *spinnaker.Client
	outputFormat string

	historyLimit int32
}

// NewCmdPipelineConfig creates the pipeline-config command.
func NewCmdPipelineConfig(ctx context.Context, client *spinnaker.Client, ioStreams *genericclioptions.IOStreams) *cobra.Command {
	pc := &pipelineConfig{
		ioStreams: ioStreams,
		client:    client,
	}

	cmd := &cobra.Command{
		Use:   "pipeline-config",
		Short: "manage the Spinnaker pipeline config.",
	}
	cmd.AddCommand(pc.get(ctx))
	cmd.AddCommand(pc.list(ctx))
	cmd.AddCommand(pc.convert(ctx))
	cmd.AddCommand(pc.history(ctx))

	return cmd
}

func (pc *pipelineConfig) get(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [application] <pipeline name>",
		Short: "Get the list or specific of an application's pipeline configurations.",
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 1); err != nil {
				return err
			}

			ctx, err = pc.client.Authenticate(ctx)
			return err
		},
		RunE: func(_ *cobra.Command, args []string) (err error) {
			application := args[0]
			logger.FromContext(ctx).Debugf("pipelineConfig.get: application: %s", application)

			var s string
			if len(args) == 2 {
				pipelineName := args[1]
				logger.FromContext(ctx).Debugf("pipelineConfig.get: pipelineName: %s", pipelineName)

				s, err = pc.client.GetPipelineConfig(ctx, application, pipelineName, pc.outputFormat)
				if err != nil {
					return err
				}
			} else {
				s, err = pc.client.GetPipelineConfigFromApplication(ctx, application, pc.outputFormat)
				if err != nil {
					return err
				}
			}

			fmt.Fprintf(pc.ioStreams.Out, s)

			return err
		},
	}

	addPrintFlags(&pc.outputFormat).AddFlags(cmd)

	return cmd
}

func (pc *pipelineConfig) list(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the all pipeline configs.",
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 0); err != nil {
				return err
			}

			ctx, err = pc.client.Authenticate(ctx)
			return err
		},
		RunE: func(*cobra.Command, []string) error {
			s, err := pc.client.ListPipelineConfigs(ctx, pc.outputFormat)
			if err != nil {
				return err
			}
			fmt.Fprintf(pc.ioStreams.Out, s)

			return nil
		},
	}

	addPrintFlags(&pc.outputFormat).AddFlags(cmd)

	return cmd
}

func (pc *pipelineConfig) history(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "history [pipeline config ID]",
		Short: "Get pipeline config history.",
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 1); err != nil {
				return err
			}

			ctx, err = pc.client.Authenticate(ctx)
			return err
		},
		RunE: func(_ *cobra.Command, args []string) error {
			pipelineConfigID := args[0]

			s, err := pc.client.GetPipelineConfigHistory(ctx, pipelineConfigID, pc.historyLimit, pc.outputFormat)
			if err != nil {
				return err
			}
			fmt.Fprintf(pc.ioStreams.Out, s)

			return nil
		},
	}

	f := cmd.Flags()
	f.Int32Var(&pc.historyLimit, "limit", 20, "Limit size of histories.")
	addPrintFlags(&pc.outputFormat).AddFlags(cmd)

	return cmd
}

func (pc *pipelineConfig) convert(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert [pipeline config ID]",
		Short: "Convert a pipeline config to a pipeline template.",
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := validateArgs(cmd, args, 1); err != nil {
				return err
			}

			ctx, err = pc.client.Authenticate(ctx)
			return err
		},
		RunE: func(_ *cobra.Command, args []string) error {
			pipelineConfigID := args[0]
			logger.FromContext(ctx).Debugf("pipelineConfig.convert: pipelineConfigID: %s", pipelineConfigID)

			s, err := pc.client.ConvertPipelineConfigToPipelineTemplate(ctx, pipelineConfigID)
			if err != nil {
				return err
			}
			fmt.Fprintf(pc.ioStreams.Out, s)

			return nil
		},
	}

	return cmd
}
