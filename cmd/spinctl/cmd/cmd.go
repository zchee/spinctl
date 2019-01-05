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
	"go.opencensus.io/stats/view"
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
	cfg := config.New()
	opts := &Options{}
	addFlags(pflags, cfg, opts)
	pflags.Parse(args)

	if !cfg.Exists() {
		if err := cfg.Create(); err != nil {
			cmd.Println(errors.WithStack(err))
		}
	}

	if path := opts.configPath; path != cfg.Path() {
		cfg.SetPath(path)
	}
	if err := cfg.Read(); err != nil {
		cmd.Println(errors.WithStack(err))
	}

	var clientOpts []spinnaker.Option
	if projectID := cfg.GoogleCloudProject; projectID != "" {
		if err := view.Register(
			ochttp.ClientSentBytesDistribution,
			ochttp.ClientReceivedBytesDistribution,
			ochttp.ClientRoundtripLatencyDistribution,
		); err != nil {
			cmd.Println(err)
		}
		clientOpts = append(clientOpts, spinnaker.WithHTTPClient(&http.Client{
			Transport: &ochttp.Transport{
				StartOptions: trace.StartOptions{
					Sampler: trace.AlwaysSample(),
				},
			},
		}))

		if err := addStackdriverExporter(ctx, projectID); err != nil {
			cmd.Println(err)
		}
	}
	client := spinnaker.NewClient(cfg, clientOpts...)

	var lv = defaultLogLevel
	if opts.debug {
		lv = zap.DebugLevel
	}
	out := cmd.OutOrStdout()
	ctx = logger.NewContext(ctx, logger.NewZapSugaredLogger(lv))

	cmd.AddCommand(NewCmdApplication(ctx, client, out))
	cmd.AddCommand(NewCmdCompletion(out))
	cmd.AddCommand(NewCmdPipeline(ctx, client, out))
	cmd.AddCommand(NewCmdPipelineTemplate(ctx, client, out))
	cmd.AddCommand(NewCmdVersion(ctx, client, out))

	return cmd
}

func addFlags(flags *pflag.FlagSet, cfg *config.Config, opts *Options) {
	flags.BoolVarP(&opts.debug, "debug", "d", false, "Use debug output")

	flags.StringVarP(&opts.configPath, "config", "c", cfg.Path(), "config file path")

	addProfilingFlags(flags)
}

func addStackdriverExporter(ctx context.Context, projectID string) error {
	stackdriverOpts := stackdriver.Options{
		ProjectID:                projectID,
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
	view.RegisterExporter(ext)

	return nil
}
