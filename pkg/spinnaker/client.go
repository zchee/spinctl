// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/zchee/spinctl/api/gate"
	"github.com/zchee/spinctl/pkg/version"
)

const (
	userAgent = "spinctl/" + version.Version
)

type Client struct {
	client *gate.APIClient
}

var defaultGateConfiguration = &gate.Configuration{
	BasePath:      "http://localhost:8084",
	DefaultHeader: make(map[string]string),
	UserAgent:     userAgent,
	HTTPClient:    http.DefaultClient,
}

type Option func(*gate.Configuration)

func WithBasePath(p string) Option {
	return func(c *gate.Configuration) {
		c.BasePath = p
	}
}

func WithHost(host string) Option {
	return func(c *gate.Configuration) {
		c.Host = host
	}
}

func WithtHeader(hdr map[string]string) Option {
	return func(c *gate.Configuration) {
		c.DefaultHeader = hdr
	}
}

func WithUserAgent(ua string) Option {
	return func(c *gate.Configuration) {
		c.UserAgent = ua
	}
}

func WithHTTPClient(hc *http.Client) Option {
	return func(c *gate.Configuration) {
		c.HTTPClient = hc
	}
}

func NewClient(opts ...Option) *Client {
	conf := defaultGateConfiguration
	for _, o := range opts {
		o(conf)
	}
	cookieJar, _ := cookiejar.New(nil)
	conf.HTTPClient.Jar = cookieJar

	return &Client{
		client: gate.NewAPIClient(conf),
	}
}
