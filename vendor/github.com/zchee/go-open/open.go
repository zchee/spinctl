// Copyright 2018 The open Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package open

import (
	"context"
	"os/exec"
)

type options struct {
	app  string
	args []string
	env  []string
	dir  string
	wait bool
}

// Command returns the cross-platform supported exec.Command.
func Command(ctx context.Context, input string, opts ...Option) (*exec.Cmd, error) {
	opt := &options{}
	for _, o := range opts {
		o(opt)
	}
	cmd, err := openContext(ctx, input, opt)
	if err != nil {
		return nil, err
	}

	if env := opt.env; env != nil {
		cmd.Env = env
	}
	if dir := opt.dir; dir != "" {
		cmd.Dir = dir
	}

	return cmd, nil
}

// Run opens the URI, files, directory or executables using the GOOS default application.
// Wait for the open command to complete.
func Run(input string, opts ...Option) error {
	return RunContext(context.Background(), input, opts...)
}

// RunContext opens the URI, files, directory or executables using the GOOS default application with context support.
// Wait for the open command to complete.
func RunContext(ctx context.Context, input string, opts ...Option) error {
	cmd, err := Command(ctx, input, opts...)
	if err != nil {
		return err
	}
	return cmd.Run()
}

// Start opens the URI, files, directory or executables using the GOOS default application.
// Don't wait for the open command to complete.
func Start(input string, opts ...Option) (func() error, error) {
	return StartContext(context.Background(), input, opts...)
}

// StartContext opens the URI, files, directory or executables using the GOOS default application with context support.
// Don't wait for the open command to complete.
func StartContext(ctx context.Context, input string, opts ...Option) (func() error, error) {
	cmd, err := Command(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	return cmd.Wait, cmd.Start()
}
