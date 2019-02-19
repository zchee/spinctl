// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commands

import (
	"context"
	"fmt"
	logpkg "log"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/profiler"
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
	"github.com/zchee/spinctl/pkg/logging"
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
		Use:                "spinctl",
		Short:              "spinctl is a command-line tool to manage Spinnaker via gate API.",
		SilenceUsage:       true,
		PersistentPreRunE:  func(*cobra.Command, []string) error { return initProfiling() },
		PersistentPostRunE: func(*cobra.Command, []string) error { return flushProfiling() },
		Version:            versionpkg.Version,
	}
	cmd.Flags().BoolP("version", "v", false, "Show "+cmd.Name()+" version.") // version flag is root only

	flags := cmd.PersistentFlags()
	cfg := config.New()
	opts := &Options{}
	addGlobalFlags(flags, cfg, opts)
	flags.Parse(args)

	var lv = atomicInfoLevel
	if opts.debug {
		lv = atomicDebugLevel
	}
	log := logging.NewZapLogger(lv)
	ctx = logging.NewContext(ctx, log)

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
	if gcpProjectID := cfg.GoogleCloudProject; gcpProjectID != "" {
		trace.ApplyConfig(trace.Config{
			DefaultSampler: trace.AlwaysSample(),
		})
		if err := view.Register(ochttp.DefaultClientViews...); err != nil {
			log.Fatal(zap.Error(err))
		}
		clientOpts = append(clientOpts, spinnaker.WithHTTPClient(&http.Client{
			Transport: &ochttp.Transport{
				StartOptions: trace.StartOptions{
					Sampler: trace.AlwaysSample(),
				},
			},
		}))
		log.Info("use ochttp.Transport")

		// Stackdriver exporter
		sdOpts := stackdriver.Options{
			ProjectID: gcpProjectID,
			OnError: func(err error) {
				log.Error(fmt.Errorf("could not log error: %v", err))
			},
			MetricPrefix: cmd.Name(),
			Context:      ctx,
		}
		sd, err := stackdriver.NewExporter(sdOpts)
		if err != nil {
			logpkg.Fatalf("failed to create stackdriver exporter: %v", err)
		}
		defer sd.Flush()
		trace.RegisterExporter(sd)
		view.RegisterExporter(sd)
		log.Info("enabled Stackdriver exporter")

		// Stackdriver Profiler
		profConf := profiler.Config{
			Service:        cmd.Name(),
			ServiceVersion: versionpkg.Version,
			MutexProfiling: true,
			ProjectID:      gcpProjectID,
		}
		if err := profiler.Start(profConf); err != nil {
			logpkg.Fatalf("failed to start stackdriver profiler: %v", err)
		}
		log.Info("enabled Stackdriver profiler")

		var span *trace.Span
		ctx, span = trace.StartSpan(ctx, "main", trace.WithSampler(trace.AlwaysSample())) // start root span
		defer span.End()
	}
	client := spinnaker.NewClient(cfg, clientOpts...)

	ioStreams := defaultIOStreams
	cmd.AddCommand(NewCmdApplication(ctx, client, ioStreams))
	cmd.AddCommand(NewCmdCompletion(ioStreams))
	cmd.AddCommand(NewCmdPipeline(ctx, client, ioStreams))
	cmd.AddCommand(NewCmdPipelineConfig(ctx, client, ioStreams))
	cmd.AddCommand(NewCmdPipelineTemplates(ctx, client, ioStreams))
	cmd.AddCommand(NewCmdVersion(ctx, client, ioStreams))

	return cmd
}

// NewStackdriverExporter apply stackdriver exporter.
func NewStackdriverExporter(ctx context.Context, projectID string) error {
	stackdriverOpts := stackdriver.Options{
		ProjectID:                projectID,
		OnError:                  func(err error) { logging.FromContext(ctx).Error(zap.Error(err)) },
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

func validateArgs(cmd *cobra.Command, args []string, expectedArgs int) error {
	use := strings.Split(cmd.Use, " ")
	if expectedArgs == 0 && len(args) > 0 {
		return errors.Errorf("%s command no needs args", use[0])
	}
	if len(args) < expectedArgs {
		use := strings.Split(cmd.Use, " ")
		return errors.Errorf("%s command required %s arg(s)", use[0], use[1:])
	}
	return nil
}
