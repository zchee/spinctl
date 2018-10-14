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

func NewCmdApplication(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "application",
		Short: "manage the Spinnaker applications.",
	}

	cmd.AddCommand(newCmdApplicationGet(ctx, client, out))
	cmd.AddCommand(newCmdApplicationList(ctx, client, out))
	cmd.AddCommand(newCmdApplicationSave(ctx, client, out))
	cmd.AddCommand(newCmdApplicationDelete(ctx, client, out))

	return cmd
}

type applicationGet struct {
	out io.Writer

	expand    bool
	outFormat string
}

func newCmdApplicationGet(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	get := &applicationGet{
		out: out,
	}

	cmd := &cobra.Command{
		Use:     "get <application name>",
		Short:   "Get the specified application.",
		Args:    cobra.ExactArgs(1),
		Example: fmt.Sprintf("  %s application get spin -x -o yaml", appName),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateArgs(args); err != nil {
				return err
			}

			name := args[0]
			logger.FromContext(ctx).Debugf("CmdApplicationGet: name: %s, expand: %t", name, get.expand)

			return client.GetApplication(ctx, get.out, name, get.expand, get.outFormat)
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&get.expand, "expand", "x", false, "expand application description.")
	f.StringVarP(&get.outFormat, "output", "o", "", "output format. One of: (json|yaml)")

	return cmd
}

type applicationList struct {
	out io.Writer

	outFormat string
}

func newCmdApplicationList(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	list := &applicationList{
		out: out,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all applications.",
		Args:    cobra.ExactArgs(0),
		Example: fmt.Sprintf("  %s application list -o yaml", appName),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateArgs(args); err != nil {
				return err
			}
			return client.ListApplications(ctx, list.out, list.outFormat)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&list.outFormat, "output", "o", "", "output format. One of: (json|yaml)")

	return cmd
}

type applicationSave struct {
	out io.Writer
}

func newCmdApplicationSave(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	_ = &applicationSave{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save the provided application.",
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}

type applicationDelete struct {
	out io.Writer
}

func newCmdApplicationDelete(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	_ = &applicationDelete{
		out: out,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete the specified application.",
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}
