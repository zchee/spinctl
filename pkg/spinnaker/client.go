// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"context"
	"net/http"
	"net/http/cookiejar"

	"github.com/pkg/errors"
	"golang.org/x/net/publicsuffix"
	"golang.org/x/oauth2"

	"github.com/zchee/spinctl/api/gate"
	"github.com/zchee/spinctl/pkg/auth"
	"github.com/zchee/spinctl/pkg/config"
	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/version"
)

const (
	userAgent = "spinctl/" + version.Version
)

type Client struct {
	client *gate.APIClient
	cfg    *config.Config
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
	// cookiejar.New return tuple secound variable always nil
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	conf.HTTPClient.Jar = cookieJar

	cfg := config.New()
	if !cfg.Exists() {
		cfg.Create()
	}
	if ep := cfg.Gate.Endpoint; ep != "" {
		conf.BasePath = ep
	}

	c := &Client{
		client: gate.NewAPIClient(conf),
		cfg:    cfg,
	}

	return c
}

func (c *Client) Authenticate(ctx context.Context) (context.Context, error) {
	authcfg := c.cfg.Auth

	if authcfg != nil && authcfg.Enable {
		tok := &oauth2.Token{}
		var err error
		if tok = authcfg.OAuth2Config.Token; tok != nil {
			conf := &oauth2.Config{
				ClientID:     authcfg.OAuth2Config.ClientID,
				ClientSecret: authcfg.OAuth2Config.ClientSecret,
				Scopes:       authcfg.OAuth2Config.Scopes,
				Endpoint: oauth2.Endpoint{
					AuthURL:  authcfg.OAuth2Config.AuthURL,
					TokenURL: authcfg.OAuth2Config.TokenURL,
				},
			}
			tokSrc := conf.TokenSource(ctx, tok)
			tok, err = tokSrc.Token()
			if err != nil {
				return nil, err
			}
		} else {
			tok, err = auth.AuthenticateOAuth2(ctx)
			if err != nil {
				return nil, err
			}
			if !tok.Valid() {
				return nil, errors.Wrapf(err, "token is invalid: %v", tok)
			}
		}
		logger.FromContext(ctx).Debugf("Authenticate: %#v", tok)

		ctx = context.WithValue(ctx, gate.ContextAccessToken, tok.AccessToken)
		authcfg.OAuth2Config.Token = tok
		if err := c.cfg.Write(); err != nil {
			return nil, err
		}
	}

	return ctx, nil
}
