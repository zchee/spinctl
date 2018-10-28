// Copyright 2018 The open Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package open

import (
	"context"
	"os/exec"
)

const (
	openCmd = "open"
)

func openContext(ctx context.Context, input string, opts *options) (*exec.Cmd, error) {
	args := []string{input}

	if optargs := opts.args; optargs != nil {
		args = append(args, append([]string{"--args"}, optargs...)...)
	}
	if app := opts.app; app != "" {
		args = append([]string{"-a", app}, args...)
	}
	if opts.wait {
		args = append([]string{"-W"}, args[0:]...)
	}

	return exec.CommandContext(ctx, openCmd, args...), nil
}
