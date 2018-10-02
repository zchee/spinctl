// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package config implements a spinctl configarations.
//
// spinctl default config filepath is respect freedesktop.org XDG Base Directory specification.
// The filepath order is:
//  $XDG_CONFIG_HOME/spinctl/config.yaml
//  $HOME/.config/spinctl/config.yaml
//  $HOME/.spinctl.yaml
package config // import "github.com/zchee/spinctl/pkg/config"
