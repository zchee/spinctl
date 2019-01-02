// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/zchee/spinctl/pkg/config"
	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
	versionpkg "github.com/zchee/spinctl/pkg/version"
)

const (
	appName = "spinctl"

	defaultLogLevel = zap.InfoLevel
)

// Options represents a root command options.
type Options struct {
	debug      bool
	configPath string
}

// NewDefaultCommand wrap of NewCommand function.
func NewDefaultCommand(ctx context.Context, args []string) *cobra.Command {
	return NewCommand(ctx, args)
}

// NewCommand creates the `spinctl` root command.
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
		Version: versionpkg.Version,
	}

	flags := cmd.PersistentFlags()
	cmd.Flags().BoolP("version", "v", false, "Show the version information.")

	conf := config.New()
	opts := &Options{}
	addFlags(flags, conf, opts)
	flags.Parse(args)

	if !conf.Exists() {
		if err := conf.Create(); err != nil {
			cmd.Println(errors.WithStack(err))
		}
	}

	if path := opts.configPath; path != conf.Path() {
		conf.SetPath(path)
	}
	if err := conf.Read(); err != nil {
		cmd.Println(errors.WithStack(err))
	}

	client := spinnaker.NewClient(conf)
	out := cmd.OutOrStdout()
	var lv = defaultLogLevel
	if opts.debug {
		lv = zap.DebugLevel
	}
	ctx = logger.NewContext(ctx, logger.NewZapSugaredLogger(lv, zapcore.AddSync(out)))

	cmd.AddCommand(NewCmdApplication(ctx, client, out))
	cmd.AddCommand(NewCmdPipeline(ctx, client, out))
	cmd.AddCommand(NewCmdVersion(ctx, client, out))

	return cmd
}

func addFlags(flags *pflag.FlagSet, conf *config.Config, opts *Options) {
	flags.BoolVarP(&opts.debug, "debug", "d", false, "Use debug output")

	flags.StringVarP(&opts.configPath, "config", "c", conf.Path(), "config file path")

	addProfilingFlags(flags)
}
