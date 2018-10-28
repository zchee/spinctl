// Copyright 2018 The open Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package open

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	winsys "golang.org/x/sys/windows"
)

const (
	cmd = "url.dll,FileProtocolHandler"
)

var (
	runDll32   = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
	withAppCmd = []string{"cmd", "/C", "start", ""}
)

func replaceInput(input string) string { return strings.NewReplacer("&", "^&").Replace(input) }

func openContext(ctx context.Context, input string, opts *options) (*exec.Cmd, error) {
	cmd := exec.CommandContext(ctx, runDll32, cmd, input)

	if optargs := opts.args; optargs != nil {
		cmd.Args = append(cmd.Args, optargs...)
	}
	if app := opts.app; app != "" {
		cmd.Args = append(withAppCmd, app, replaceInput(input))
	}
	if opts.wait {
		cmd.SysProcAttr = &winsys.SysProcAttr{HideWindow: true}
	}

	return cmd, nil
}
