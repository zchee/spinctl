// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/zchee/spinctl/pkg/spinnaker"
)

type application struct {
	client *spinnaker.Client
}

func NewCmdApplication(ctx context.Context, client *spinnaker.Client) *cobra.Command {
	a := &application{
		client: client,
	}

	cmd := &cobra.Command{
		Use:   "application",
		Short: "manage the Spinnaker applications.",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "Get the specified application.",
		RunE: func(*cobra.Command, []string) error {
			return a.runGet(ctx)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all applications.",
		RunE: func(*cobra.Command, []string) error {
			return a.runList(ctx)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "save",
		Short: "Save the provided application.",
		RunE: func(*cobra.Command, []string) error {
			return a.runSave(ctx)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "delete",
		Short: "Delete the specified application.",
		RunE: func(*cobra.Command, []string) error {
			return a.runDelete(ctx)
		},
	})

	return cmd
}

func (a *application) init(ctx context.Context, cmd *cobra.Command, args []string) error {
	return nil
}

func (a *application) runGet(ctx context.Context) error {
	return nil
}

func (a *application) runList(ctx context.Context) error {
	return a.client.ListApplications(ctx)
}

func (a *application) runSave(ctx context.Context) error {
	return nil
}

func (a *application) runDelete(ctx context.Context) error {
	return nil
}
