// Copyright 2019 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/spf13/pflag"
	"github.com/zchee/spinctl/pkg/config"
)

func NewPrintFlags() *genericclioptions.PrintFlags {
	outputFormat := ""
	templateArgPtr := ""

	return &genericclioptions.PrintFlags{
		OutputFormat:       &outputFormat,
		JSONYamlPrintFlags: genericclioptions.NewJSONYamlPrintFlags(),
		TemplatePrinterFlags: &genericclioptions.KubeTemplatePrintFlags{
			GoTemplatePrintFlags: &genericclioptions.GoTemplatePrintFlags{
				TemplateArgument: &templateArgPtr,
			},
			JSONPathPrintFlags: &genericclioptions.JSONPathPrintFlags{
				TemplateArgument: &templateArgPtr,
			},
			TemplateArgument: &templateArgPtr,
		},
	}
}

func addGlobalFlags(flags *pflag.FlagSet, cfg *config.Config, opts *Options) {
	flags.BoolVarP(&opts.debug, "debug", "d", false, "Use debug output")

	flags.StringVarP(&opts.configPath, "config", "c", cfg.Path(), "config file path")

	addProfilingFlags(flags)
}

var (
	profileName   string
	profileOutput string
)

func initProfiling() error {
	switch profileName {
	case "none":
		return nil
	case "cpu":
		f, err := os.Create(profileOutput)
		if err != nil {
			return err
		}
		return pprof.StartCPUProfile(f)
	// Block and mutex profiles need a call to Set{Block,Mutex}ProfileRate to
	// output anything. We choose to sample all events.
	case "block":
		runtime.SetBlockProfileRate(1)
		return nil
	case "mutex":
		runtime.SetMutexProfileFraction(1)
		return nil
	default:
		// Check the profile name is valid.
		if profile := pprof.Lookup(profileName); profile == nil {
			return fmt.Errorf("unknown profile '%s'", profileName)
		}
	}

	return nil
}

func flushProfiling() error {
	switch profileName {
	case "none":
		return nil
	case "cpu":
		pprof.StopCPUProfile()
	case "heap":
		runtime.GC()
		fallthrough
	default:
		profile := pprof.Lookup(profileName)
		if profile == nil {
			return nil
		}
		f, err := os.Create(profileOutput)
		if err != nil {
			return err
		}
		profile.WriteTo(f, 0)
	}

	return nil
}

func addProfilingFlags(flags *pflag.FlagSet) {
	flags.StringVar(&profileName, "profile", "none", "Name of profile to capture. One of (none|cpu|heap|goroutine|threadcreate|block|mutex)")
	flags.MarkHidden("profile")
	flags.StringVar(&profileOutput, "profile-output", "profile.pprof", "Name of the file to write the profile to")
	flags.MarkHidden("profile-output")
}
