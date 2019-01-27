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

type pipelineConfig struct {
	out    io.Writer
	client *spinnaker.Client
	output string

	historyLimit int32
}

// NewCmdPipelineConfig creates the pipeline-config command.
func NewCmdPipelineConfig(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	pc := &pipelineConfig{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   "pipeline-config",
		Short: "manage the Spinnaker pipeline config.",
	}
	cmd.AddCommand(pc.get(ctx))
	cmd.AddCommand(pc.list(ctx))
	cmd.AddCommand(pc.convert(ctx))

	return cmd
}

func (pc *pipelineConfig) get(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [application] <pipeline name>",
		Short: "Get the list or specific of an application's pipeline configurations.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = pc.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// TODO(zchee): validate arg length.
			application := args[0]
			logger.FromContext(ctx).Debugf("pipelineConfig.get: application: %s", application)

			var s string
			if len(args) == 2 {
				pipelineName := args[1]
				logger.FromContext(ctx).Debugf("pipelineConfig.get: pipelineName: %s", pipelineName)

				s, err = pc.client.GetPipelineConfig(ctx, application, pipelineName, pc.output)
				if err != nil {
					return err
				}
			} else {
				s, err = pc.client.GetPipelineConfigFromApplication(ctx, application, pc.output)
				if err != nil {
					return err
				}
			}

			fmt.Fprintf(pc.out, s)

			return err
		},
	}

	f := cmd.Flags()
	f.StringVarP(&pc.output, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

func (pc *pipelineConfig) list(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the all pipeline configs.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = pc.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(*cobra.Command, []string) error {
			// TODO(zchee): validate arg length.
			s, err := pc.client.ListPipelineConfigs(ctx, pc.output)
			if err != nil {
				return err
			}
			fmt.Fprintf(pc.out, s)

			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&pc.output, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

func (pc *pipelineConfig) history(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "history [pipeline config ID]",
		Short: "Get pipeline config history.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = pc.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(_ *cobra.Command, args []string) error {
			// TODO(zchee): validate arg length.
			pipelineConfigID := args[0]

			s, err := pc.client.GetPipelineConfigHistory(ctx, pipelineConfigID, pc.historyLimit, pc.output)
			if err != nil {
				return err
			}
			fmt.Fprintf(pc.out, s)

			return nil
		},
	}

	f := cmd.Flags()
	f.Int32Var(&pc.historyLimit, "limit", 0, "Limit size of histories.")
	f.StringVarP(&pc.output, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

func (pc *pipelineConfig) convert(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert [pipeline config id]",
		Short: "Convert a pipeline config to a pipeline template.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = pc.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO(zchee): validate arg length
			pipelineConfigID := args[0]
			logger.FromContext(ctx).Debugf("pipelineConfig.convert: pipelineConfigID: %s", pipelineConfigID)

			s, err := pc.client.ConvertPipelineConfigToPipelineTemplate(ctx, pipelineConfigID)
			if err != nil {
				return err
			}
			fmt.Fprintf(pc.out, s)

			return nil
		},
	}

	return cmd
}
