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

	expand bool
}

func NewCmdApplication(ctx context.Context, client *spinnaker.Client) *cobra.Command {
	a := &application{
		client: client,
		log:    logger.FromContext(ctx),
	}

	cmd := &cobra.Command{
		Use:   "application",
		Short: "manage the Spinnaker applications.",
		PersistentPreRunE: func(*cobra.Command, []string) (err error) {
			ctx, err = a.client.Authenticate(ctx)
			return err
		},
	}

	flags := cmd.PersistentFlags()
	flags.BoolVarP(&a.expand, "expand", "x", false, "expand application.")

	cmd.AddCommand(&cobra.Command{
		Use:   "get <application name>",
		Short: "Get the specified application.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.ValidateArgs(args)

			name := args[0]
			a.log.Debug("cmdGet",
				zap.String("name", name),
				zap.Bool("a.expand", a.expand),
			)

			return a.client.GetApplication(ctx, name, a.expand)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all applications.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cobra.NoArgs(cmd, args); err != nil {
				return err
			}

			return a.client.ListApplications(ctx)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "save",
		Short: "Save the provided application.",
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "delete",
		Short: "Delete the specified application.",
		RunE: func(*cobra.Command, []string) error {
			return errors.New("not implements yet")
		},
	})

	return cmd
}
