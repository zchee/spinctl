// Copyright 2019 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

// BasicConfig represents a basic authentication data.
type BasicConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// IsValid reports whether the BasicConfig is valid.
func (b *BasicConfig) IsValid() bool {
	return b.Username != "" && b.Password != ""
}
