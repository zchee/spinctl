// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"net/http"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	apioption "google.golang.org/api/option"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/zchee/spinctl/pkg/config"
	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/spinnaker"
	versionpkg "github.com/zchee/spinctl/pkg/version"
)

var (
	defaultIOStreams = &genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}

	atomicInfoLevel  = zap.NewAtomicLevelAt(zap.InfoLevel)
	atomicDebugLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
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
	flags.BoolP("version", "v", false, "Show "+cmd.Name()+" version.") // version flag is root only

	persistentFlags := cmd.PersistentFlags()
	cfg := config.New()
	opts := &Options{}
	addGlobalFlags(persistentFlags, cfg, opts)
	printFlags := NewPrintFlags()
	printFlags.AddFlags(cmd)

	persistentFlags.Parse(args)

	var lv = atomicInfoLevel
	if opts.debug {
		lv = atomicDebugLevel
	}
	log := logger.NewZapLogger(lv)
	ctx = logger.NewContext(ctx, log)

	if !cfg.Exists() {
		if err := cfg.Create(); err != nil {
			log.Fatal(zap.Error(err))
		}
	}

	if path := opts.configPath; path != cfg.Path() {
		cfg.SetPath(path)
	}
	if err := cfg.Read(); err != nil {
		log.Fatal(zap.Error(err))
	}

	var clientOpts []spinnaker.Option
	if projectID := cfg.GoogleCloudProject; projectID != "" {
		log.Debug("use StackdriverExporter")
		if err := view.Register(
			ochttp.ClientSentBytesDistribution,
			ochttp.ClientReceivedBytesDistribution,
			ochttp.ClientRoundtripLatencyDistribution,
		); err != nil {
			log.Fatal(zap.Error(err))
		}
		clientOpts = append(clientOpts, spinnaker.WithHTTPClient(&http.Client{
			Transport: &ochttp.Transport{
				StartOptions: trace.StartOptions{
					Sampler: trace.AlwaysSample(),
				},
			},
		}))

		if err := NewStackdriverExporter(ctx, projectID); err != nil {
			log.Fatal(zap.Error(err))
		}
	}
	client := spinnaker.NewClient(cfg, clientOpts...)

	ioStreams := defaultIOStreams
	cmd.AddCommand(NewCmdApplication(ctx, client, ioStreams))
	cmd.AddCommand(NewCmdCompletion(ioStreams))
	cmd.AddCommand(NewCmdPipeline(ctx, client, ioStreams))
	cmd.AddCommand(NewCmdPipelineConfig(ctx, client, ioStreams))
	cmd.AddCommand(NewCmdPipelineTemplate(ctx, client, ioStreams))
	cmd.AddCommand(NewCmdVersion(ctx, client, ioStreams))

	return cmd
}

func NewStackdriverExporter(ctx context.Context, projectID string) error {
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
