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

type application struct {
	out    io.Writer
	client *spinnaker.Client
	output string

	getExpand bool

	listAccount string
	listOwner   string
}

// NewCmdApplication creates the application command.
func NewCmdApplication(ctx context.Context, client *spinnaker.Client, out io.Writer) *cobra.Command {
	a := &application{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:     "application",
		Aliases: []string{"app"},
		Short:   "manage the Spinnaker applications.",
	}

	cmd.AddCommand(a.get(ctx))
	cmd.AddCommand(a.list(ctx))
	cmd.AddCommand(a.save(ctx))
	cmd.AddCommand(a.delete(ctx))

	return cmd
}

func (a *application) get(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get <application name>",
		Short:   "Get the specified application.",
		Args:    cobra.ExactArgs(1),
		Example: "  spinctl application get spin -x -o yaml",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = a.client.Authenticate(ctx)
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
			logger.FromContext(ctx).Debugf("CmdApplicationGet: name: %s, expand: %t", name, a.getExpand)

			s, err := a.client.GetApplication(ctx, name, a.getExpand, a.output)
			if err != nil {
				return err
			}
			fmt.Fprintf(a.out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&a.getExpand, "expand", "x", false, "expand application description.")
	f.StringVarP(&a.output, "output", "o", "", "output format. One of: (json|yaml)")

	return cmd
}

func (a *application) list(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all applications.",
		Args:    cobra.ExactArgs(0),
		Example: "  spinctl application list -o yaml",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = a.client.Authenticate(ctx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := a.client.ListApplications(ctx, a.listAccount, a.listOwner, a.output)
			if err != nil {
				return err
			}
			fmt.Fprintf(a.out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVar(&a.listAccount, "account", "", "filters results to only include applications deployed in the specified account")
	f.StringVar(&a.listOwner, "owner", "", "filteres results to only include applications owned by the specified email")
	f.StringVarP(&a.output, "output", "o", "", "output format. One of: (json|yaml)")

	return cmd
}

func (a *application) save(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save the provided application.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = a.client.Authenticate(ctx)
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

func (a *application) delete(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete the specified application.",
		PreRunE: func(*cobra.Command, []string) error {
			var err error
			ctx, err = a.client.Authenticate(ctx)
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
