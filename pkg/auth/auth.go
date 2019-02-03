// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

// Config implements a authentication config for Spinnaker.
type Config struct {
	Enable bool          `yaml:"enable,omitempty"`
	Basic  *BasicConfig  `yaml:"basic,omitempty"`
	OAuth2 *OAuth2Config `yaml:"oauth2,omitempty"`
}
