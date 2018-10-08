// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap/zapcore"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

const (
	appName = "spinctl"
)

func NewDefaultCommand(ctx context.Context, args []string) *cobra.Command {
	return NewCommand(ctx, args)
}

func NewCommand(ctx context.Context, args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          appName,
		Short:        fmt.Sprintf("%s is a command-line tool to manage Spinnaker via gate API.", appName),
		SilenceUsage: true,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return initProfiling()
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			return flushProfiling()
		},
	}

	out := cmd.OutOrStdout()

	ctx = logger.NewContext(ctx, logger.NewZapSugaredLogger(zapcore.InfoLevel, zapcore.AddSync(out)))

	flags := cmd.PersistentFlags()
	addProfilingFlags(flags)
	flags.Parse(args)

	client := spinnaker.NewClient()
	ctx, err := client.Authenticate(ctx)
	if err != nil {
		logger.FromContext(ctx).Fatalf("%v\n", err)
	}

	cmd.AddCommand(NewCmdApplication(ctx, client, out))

	return cmd
}
