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
	client *spinnaker.Client
	out    io.Writer

	expand    bool
	outFormat string
}

func newCmdPipelineGet(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	_ = &pipelineGet{
		client: client,
		out:    out,
	}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the specified pipeline.",
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}

type pipelineList struct {
	client *spinnaker.Client
	out    io.Writer

	outFormat string
}

func newCmdPipelineList(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	list := &pipelineList{
		client: client,
		out:    out,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List the pipelines for the provided application.",
		Args:    cobra.ExactArgs(1),
		Example: fmt.Sprintf("  %s pipeline list <application name> -o yaml", appName),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateArgs(args); err != nil {
				return err
			}

			name := args[0]
			logger.FromContext(ctx).Debugf("CmdPipelineList: name: %s", name)

			return list.client.ListPipelines(ctx, list.out, name, list.outFormat)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&list.outFormat, "output", "o", "", "Output format. One of: (json|yaml)")

	return cmd
}

type pipelineSave struct {
	client *spinnaker.Client
	out    io.Writer
}

func newCmdPipelineSave(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	_ = &pipelineSave{
		client: client,
		out:    out,
	}

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save the provided pipeline.",
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}

type pipelineDelete struct {
	client *spinnaker.Client
	out    io.Writer
}

func newCmdPipelineDelete(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	_ = &pipelineDelete{
		client: client,
		out:    out,
	}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete the provided pipeline.",
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}

type pipelineExecute struct {
	client *spinnaker.Client
	out    io.Writer
}

func newCmdPipelineExecute(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	_ = &pipelineExecute{
		client: client,
		out:    out,
	}

	cmd := &cobra.Command{
		Use:   "execute",
		Short: "Execute the provided pipeline.",
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}
