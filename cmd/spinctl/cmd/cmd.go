// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	apioption "google.golang.org/api/option"

	"github.com/zchee/spinctl/pkg/config"
	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
	versionpkg "github.com/zchee/spinctl/pkg/version"
)

const (
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
		Use:          "spinctl",
		Short:        "spinctl is a command-line tool to manage Spinnaker via gate API.",
		SilenceUsage: true,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return initProfiling()
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			return flushProfiling()
		},
		Version: versionpkg.Version,
	}

	flags := cmd.Flags()
	flags.BoolP("version", "v", false, "Show the version information.") // version flag is root only

	pflags := cmd.PersistentFlags()
	conf := config.New()
	opts := &Options{}
	addFlags(pflags, conf, opts)
	pflags.Parse(args)

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

	hc := &http.Client{
		Transport: &ochttp.Transport{},
	}
	client := spinnaker.NewClient(conf, spinnaker.WithHTTPClient(hc))

	var lv = defaultLogLevel
	if opts.debug {
		lv = zap.DebugLevel
	}
	out := cmd.OutOrStdout()
	ctx = logger.NewContext(ctx, logger.NewZapSugaredLogger(lv))

	if err := addStackdriverExporter(ctx); err != nil {
		cmd.Println(err)
	}

	cmd.AddCommand(NewCmdApplication(ctx, client, out))
	cmd.AddCommand(NewCmdCompletion(out))
	cmd.AddCommand(NewCmdPipeline(ctx, client, out))
	cmd.AddCommand(NewCmdPipelineTemplate(ctx, client, out))
	cmd.AddCommand(NewCmdVersion(ctx, client, out))

	return cmd
}

func addFlags(flags *pflag.FlagSet, conf *config.Config, opts *Options) {
	flags.BoolVarP(&opts.debug, "debug", "d", false, "Use debug output")

	flags.StringVarP(&opts.configPath, "config", "c", conf.Path(), "config file path")

	addProfilingFlags(flags)
}

func addStackdriverExporter(ctx context.Context) error {
	stackdriverOpts := stackdriver.Options{
		ProjectID:                "xxx",
		OnError:                  func(err error) { logger.FromContext(ctx).Error(zap.Error(err)) },
		MonitoringClientOptions:  []apioption.ClientOption{},
		TraceClientOptions:       []apioption.ClientOption{},
		BundleCountThreshold:     0,
		TraceSpansBufferMaxBytes: 8 * 1024 * 1024, // 8MB
		MetricPrefix:             "spinctl/",
		Context:                  ctx,
		Timeout:                  5 * time.Second,
	}
	ext, err := stackdriver.NewExporter(stackdriverOpts)
	if err != nil {
		return errors.WithStack(err)
	}
	trace.RegisterExporter(ext)

	return nil
}
