// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

type application struct {
	ioStreams    *genericclioptions.IOStreams
	client       *spinnaker.Client
	outputFormat string

	getExpand bool

	listAccount string
	listOwner   string
}

// NewCmdApplication creates the application command.
func NewCmdApplication(ctx context.Context, client *spinnaker.Client, ioStreams *genericclioptions.IOStreams) *cobra.Command {
	a := &application{
		ioStreams: ioStreams,
		client:    client,
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
		PreRunE: func(*cobra.Command, []string) (err error) {
			ctx, err = a.client.Authenticate(ctx)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateArgs(args); err != nil {
				return err
			}

			name := args[0]
			logger.FromContext(ctx).Debugf("application.get: name: %s, expand: %t", name, a.getExpand)
			logger.FromContext(ctx).Debug(a.outputFormat)

			s, err := a.client.GetApplication(ctx, name, a.getExpand, a.outputFormat)
			if err != nil {
				return err
			}
			fmt.Fprintf(a.ioStreams.Out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&a.getExpand, "expand", "x", false, "expand application description.")
	addPrintFlags(&a.outputFormat).AddFlags(cmd)

	return cmd
}

func (a *application) list(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all applications.",
		Args:    cobra.ExactArgs(0),
		Example: "  spinctl application list -o yaml",
		PreRunE: func(*cobra.Command, []string) (err error) {
			ctx, err = a.client.Authenticate(ctx)
			return err
		},
		RunE: func(*cobra.Command, []string) error {
			s, err := a.client.ListApplications(ctx, a.listAccount, a.listOwner, a.outputFormat)
			if err != nil {
				return err
			}
			fmt.Fprintf(a.ioStreams.Out, s)
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVar(&a.listAccount, "account", "", "filters results to only include applications deployed in the specified account")
	f.StringVar(&a.listOwner, "owner", "", "filteres results to only include applications owned by the specified email")
	addPrintFlags(&a.outputFormat).AddFlags(cmd)

	return cmd
}

func (a *application) save(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save the provided application.",
		PreRunE: func(*cobra.Command, []string) (err error) {
			ctx, err = a.client.Authenticate(ctx)
			return err
		},
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}

func (a *application) delete(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete the specified application.",
		PreRunE: func(*cobra.Command, []string) (err error) {
			ctx, err = a.client.Authenticate(ctx)
			return err
		},
		RunE: func(*cobra.Command, []string) error {
			return errNotImplementedYet
		},
	}

	return cmd
}
