// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command spinctl is a command-line tool to manage Spinnaker via gate API.
package main

import (
	"context"
	"os"

	"github.com/zchee/spinctl/pkg/commands"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := commands.NewDefaultCommand(ctx, os.Args[1:])
	if err := c.Execute(); err != nil {
		os.Exit(1) //nolint:gocritic
	}
}
