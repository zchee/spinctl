// Copyright 2018 The open Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux android

package open

import (
	"context"
	"fmt"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"

	sys "golang.org/x/sys/unix"
)

const (
	openCmd = "xdg-open"
)

// ref: xdg-open
//  https://sources.debian.org/src/xdg-utils/1.1.3-1/scripts/xdg-open.in/
func openContext(ctx context.Context, input string, opts *options) (*exec.Cmd, error) {
	args := []string{input}

	if optargs := opts.args; optargs != nil {
		args = append(args, optargs...)
	}

	cmd := exec.CommandContext(ctx, openCmd, args...)

	if app := opts.app; app != "" {
		cmd.Args = append([]string{app}, cmd.Args...)
	}

	if !opts.wait {
		u, err := user.Current()
		if err != nil {
			return nil, fmt.Errorf("error getting current user: %v", err)
		}
		uid, err := strconv.Atoi(u.Uid)
		if err != nil {
			return nil, fmt.Errorf("error converting Uid=%s to integer: %v", u.Uid, err)
		}
		gid, err := strconv.Atoi(u.Gid)
		if err != nil {
			return nil, fmt.Errorf("error converting Gid=%s to integer: %v", u.Gid, err)
		}
		cmd.SysProcAttr = &sys.SysProcAttr{
			Credential: &syscall.Credential{
				Uid:         uint32(uid),
				Gid:         uint32(gid),
				NoSetGroups: true,
			},
			Cloneflags: sys.CLONE_DETACHED,
			Foreground: false,
		}
	}

	return cmd, nil
}
