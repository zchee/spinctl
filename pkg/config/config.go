// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
	xdgbasedir "github.com/zchee/go-xdgbasedir"
	"github.com/zchee/go-xdgbasedir/home"
	yaml "gopkg.in/yaml.v2"

	"github.com/zchee/spinctl/pkg/auth"
)

var (
	configPath string
)

// init sets xdgbasedir Mode to Unix for darwin GOOS.
func init() {
	xdgbasedir.Mode = xdgbasedir.Unix

	configDirBase := xdgbasedir.ConfigHome()
	switch {
	case isDirExist(configDirBase):
		configPath = filepath.Join(configDirBase, "spinctl", "config.yaml")
	default:
		configPath = filepath.Join(home.Dir(), ".spinctl.yaml")
	}
}

// Config implements a spinctl configarations.
type Config struct {
	Gate Gate         `yaml:"gate,omitempty"`
	Auth *auth.Config `yaml:"auth,omitempty"`

	path          string
	exists        bool
	dirCreateOnce sync.Once
}

// Gate implements a Spinnaker gate API configarations.
type Gate struct {
	Endpoint string `yaml:"endpoint,omitempty"`
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
		Auth: new(auth.Config),
		path: configPath,
	}

	return c
}

// Path returns the config filepath.
func (c *Config) Path() string {
	return c.path
}

// Exists returns whether the exist config file.
func (c *Config) Exists() bool {
	return c.exists
}

// Create creates the config file.
func (c *Config) Create() (err error) {
	c.dirCreateOnce.Do(func() {
		configDir := filepath.Dir(c.path)
		// mkdir with secure permission if not exist
		if _, statErr := os.Stat(configDir); os.IsNotExist(statErr) {
			if e := os.MkdirAll(configDir, 0700); e != nil {
				err = e
			}
		}

		// create config file(c.path) with secure permission if not exist
		if _, statErr := os.Stat(c.path); statErr != nil && os.IsNotExist(statErr) {
			if _, e := os.OpenFile(c.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600); e != nil {
				err = e
			}
		}

		c.exists = true
	})
	if err != nil {
		return errors.Wrapf(err, "failed to create config file to %s", c.path)
	}

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
func isDirExist(dir string) bool {
	d, err := os.Stat(dir)
	if err == nil && d.IsDir() {
		return true
	}

	return false
}
