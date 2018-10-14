// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap/zapcore"

	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
	"github.com/zchee/spinctl/pkg/version"
)

const (
	appName = "spinctl"

	defaultLogLevel = zapcore.InfoLevel
)

type Options struct {
	debug bool
}

func NewDefaultCommand(ctx context.Context, args []string) *cobra.Command {
	return NewCommand(ctx, args)
}

func NewCommand(ctx context.Context, args []string) *cobra.Command {
	opts := &Options{}

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
		Version: version.Version,
	}

	flags := cmd.PersistentFlags()
	addFlags(flags, opts)
	flags.Parse(args)

	out := cmd.OutOrStdout()

	lv := defaultLogLevel
	if opts.debug {
		lv = zapcore.DebugLevel
	}

	ctx = logger.NewContext(ctx, logger.NewZapSugaredLogger(lv, zapcore.AddSync(out)))
	client := spinnaker.NewClient()

	cmd.AddCommand(NewCmdApplication(ctx, client, out))
	cmd.AddCommand(NewCmdPipeline(ctx, client, out))

	return cmd
}

func addFlags(flags *pflag.FlagSet, opts *Options) {
	flags.BoolVarP(&opts.debug, "debug", "d", false, "Use debug output")
	flags.BoolP("version", "v", false, "Show the version information.")

	addProfilingFlags(flags)
}
