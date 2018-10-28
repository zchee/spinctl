// Copyright 2018 The open Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package open

// Option represents a functional pattern option type for options.
type Option func(*options)

// WithApplication sets the open with specific application to open command.
func WithApplication(app string) Option {
	return func(o *options) {
		o.app = app
	}
}

// WithArgs sets the any arguments to open command.
func WithArgs(args []string) Option {
	return func(o *options) {
		o.args = args
	}
}

// WithEnv sets the environment variables to open command.
func WithEnv(env []string) Option {
	return func(o *options) {
		o.env = env
	}
}

// WithDir sets the directory to open command.
func WithDir(dir string) Option {
	return func(o *options) {
		o.dir = dir
	}
}

// WithWait wait for the open command to complete.
func WithWait() Option {
	return func(o *options) {
		o.wait = true
	}
}
