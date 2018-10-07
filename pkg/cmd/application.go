// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

type application struct {
	client *spinnaker.Client
	log    *zap.Logger
}

func NewCmdApplication(ctx context.Context, client *spinnaker.Client) *cobra.Command {
	a := &application{
		client: client,
		log:    logger.FromContext(ctx),
	}

	cmd := &cobra.Command{
		Use:   "application",
		Short: "manage the Spinnaker applications.",
	}

	cmd.AddCommand(a.newCmdApplicationGet(ctx))
	cmd.AddCommand(a.newCmdApplicationList(ctx))
	cmd.AddCommand(a.newCmdApplicationSave(ctx))
	cmd.AddCommand(a.newCmdApplicationDelete(ctx))

	return cmd
}

type applicationGet struct {
	expand bool
}

func (a *application) newCmdApplicationGet(ctx context.Context) *cobra.Command {
	get := &applicationGet{}

	cmd := &cobra.Command{
		Use:   "get <application name>",
		Short: "Get the specified application.",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(*cobra.Command, []string) (err error) {
			ctx, err = a.client.Authenticate(ctx)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.ValidateArgs(args)

			name := args[0]
			a.log.Debug("CmdApplicationGet",
				zap.String("name", name),
				zap.Bool("expand", get.expand),
			)

			return a.client.GetApplication(ctx, name, get.expand)
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&get.expand, "expand", "x", false, "expand application description.")

	return cmd
}

type applicationList struct{}

func (a *application) newCmdApplicationList(ctx context.Context) *cobra.Command {
	_ = &applicationList{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all applications.",
		PreRunE: func(*cobra.Command, []string) (err error) {
			ctx, err = a.client.Authenticate(ctx)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cobra.NoArgs(cmd, args); err != nil {
				return err
			}

			return a.client.ListApplications(ctx)
		},
	}

	return cmd
}

type applicationSave struct{}

func (a *application) newCmdApplicationSave(ctx context.Context) *cobra.Command {
	_ = &applicationSave{}

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save the provided application.",
		PreRunE: func(*cobra.Command, []string) (err error) {
			ctx, err = a.client.Authenticate(ctx)
			return err
		},
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}

type applicationDelete struct{}

func (a *application) newCmdApplicationDelete(ctx context.Context) *cobra.Command {
	_ = &applicationDelete{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete the specified application.",
		PreRunE: func(*cobra.Command, []string) (err error) {
			ctx, err = a.client.Authenticate(ctx)
			return err
		},
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	}

	return cmd
}
