// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command spinctl is a command-line tool to manage Spinnaker via gate API.
package main

import (
	"context"
	"os"

	"go.uber.org/zap/zapcore"

	"github.com/zchee/spinctl/pkg/cmd"
	"github.com/zchee/spinctl/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.NewZapSugaredLogger(zapcore.InfoLevel, os.Stdout)
	ctx = logger.NewContext(ctx, log)

	c := cmd.NewDefaultCommand(ctx, os.Args[1:])
	if err := c.Execute(); err != nil {
		switch err.(type) {
		default:
			log.Fatal("main")
		}
	}
}
