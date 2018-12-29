// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

// NewCmdPipeline creates the pipeline command.
func NewCmdPipeline(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pipeline",
		Short: "manage the Spinnaker pipelines.",
	}

	cmd.AddCommand(newCmdPipelineGet(ctx, client, out))
	cmd.AddCommand(newCmdPipelineList(ctx, client, out))
	cmd.AddCommand(newCmdPipelineSave(ctx, client, out))
	cmd.AddCommand(newCmdPipelineDelete(ctx, client, out))
	cmd.AddCommand(newCmdPipelineExecute(ctx, client, out))

	return cmd
}

type pipelineGet struct {
	out io.Writer

	outFormat string
}

func newCmdPipelineGet(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	get := &pipelineGet{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the specified pipeline.",
		Args:  cobra.ExactArgs(2),
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

			application := args[0]
			pipelineName := args[1]
			logger.FromContext(ctx).Debugf("CmdPipelineGet: application: %s, pipelineName: %s", application, pipelineName)

			return client.GetPipeline(ctx, get.out, application, pipelineName, get.outFormat)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&get.outFormat, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

type pipelineList struct {
	out io.Writer

	outFormat string
}

func newCmdPipelineList(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	list := &pipelineList{
		out: out,
	}

	cmd := &cobra.Command{
		Use:     "list <application name>",
		Aliases: []string{"ls"},
		Short:   "List the pipelines for the provided application.",
		Args:    cobra.ExactArgs(1),
		Example: fmt.Sprintf("  %s pipeline list spin -o yaml", appName),
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

			name := args[0]
			logger.FromContext(ctx).Debugf("CmdPipelineList: name: %s", name)

			return client.ListPipelines(ctx, list.out, name, list.outFormat)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&list.outFormat, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

type pipelineSave struct {
	out io.Writer
}

func newCmdPipelineSave(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	_ = &pipelineSave{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save the provided pipeline.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}

type pipelineDelete struct {
	out io.Writer
}

func newCmdPipelineDelete(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	_ = &pipelineDelete{
		out: out,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete the provided pipeline.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}

type pipelineExecute struct {
	out io.Writer
}

func newCmdPipelineExecute(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	_ = &pipelineExecute{
		out: out,
	}

	cmd := &cobra.Command{
		Use:     "execute",
		Aliases: []string{"exec"},
		Short:   "Execute the provided pipeline.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}
