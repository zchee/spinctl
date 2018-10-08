// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	xdgbasedir "github.com/zchee/go-xdgbasedir"
	"github.com/zchee/go-xdgbasedir/home"
	yaml "gopkg.in/yaml.v2"

	"github.com/zchee/spinctl/pkg/auth"
)

// Config implements a spinctl configarations.
type Config struct {
	Gate Gate         `yaml:"gate,omitempty"`
	Auth *auth.Config `yaml:"auth,omitempty"`

	path   string
	exists bool
}

// Gate implements a Spinnaker gate API configarations.
type Gate struct {
	Endpoint string `yaml:"endpoint,omitempty"`
}

// init sets xdgbasedir Mode to Unix for darwin GOOS.
func init() {
	xdgbasedir.Mode = xdgbasedir.Unix
}

// New loads the spinctl config file and returns the new Config.
// If config file is not exists, will create the directory and file.
//
// The path order is:
//  $XDG_CONFIG_HOME/spinctl/config.yaml
//  $HOME/.config/spinctl/config.yaml
//  $HOME/.spinctl.yaml
func New() *Config {
	c := &Config{
		Auth: &auth.Config{
			OAuth2Config: new(auth.OAuth2Config),
		},
		path: filepath.Join(xdgbasedir.ConfigHome(), "spinctl", "config.yaml"),
	}

	if dir := isDirExist(xdgbasedir.ConfigHome()); dir == "" {
		// fallback to use $HOME directly
		c.path = filepath.Join(home.Dir(), ".spinctl.yaml")
	}

	if _, err := os.Stat(c.path); err == nil {
		if err := c.Read(); err != nil {
			return c
		}
		c.exists = true
	}

	return c
}

// Exists returns whether the exist config file.
func (c *Config) Exists() bool {
	return c.exists
}

// Create creates the config file.
func (c *Config) Create() error {
	dir := filepath.Dir(c.path)
	// mkdir with secure permission if not exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return err
		}
	}

	// create config file with secure permission if not exist
	if _, err := os.Stat(c.path); err != nil && os.IsNotExist(err) {
		if _, err := os.OpenFile(c.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600); err != nil {
			return err
		}
	}

	c.exists = true
	return nil
}

// Read reads the config file to c.
func (c *Config) Read() error {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

// Write writes the c to config file.
func (c *Config) Write() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(c.path, data, 0600); err != nil {
		return err
	}

	return nil
}

// isDirExist checks whether the dir exists and which is directory.
func isDirExist(dir string) string {
	d, err := os.Stat(dir)
	if err == nil && d.IsDir() {
		return dir
	}
	return ""
}
