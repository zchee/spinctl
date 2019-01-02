// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package version provides the spinctl version.
package version

var (
	// version indicates which version of the binary is running.
	version = "dev"

	// GitCommit indicates which git hash the binary was built off of.
	gitCommit = ""

	// Version is the current spinctl version.
	Version = version + "@" + gitCommit
)
