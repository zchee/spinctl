// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"io"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap/zapcore"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
)

func NewDefaultCommand(ctx context.Context, args []string) *cobra.Command {
	return NewCommand(ctx, os.Stdin, os.Stdout, os.Stderr, args)
}

func NewCommand(ctx context.Context, in io.Reader, out, er io.Writer, args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "spinctl",
		Short:        "spinctl is a command-line tool to manage Spinnaker via gate API.",
		SilenceUsage: true,
	}

	flags := cmd.PersistentFlags()
	flags.Parse(args)

	ctx = logger.NewContext(ctx, logger.NewZapLogger(zapcore.DebugLevel))

	client := spinnaker.NewClient()
	cmd.AddCommand(NewCmdApplication(ctx, client))

	return cmd
}
