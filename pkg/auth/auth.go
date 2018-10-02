// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

// AuthConfig implements a authentication config for Spinnaker.
type AuthConfig struct {
	Enable        bool `yaml:"enable"`
	*OAuth2Config `yaml:"oauth2,omitempty"`
}
