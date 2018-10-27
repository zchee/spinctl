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

var (
	userAgent = "spinctl/" + version.Version
)

type Client struct {
	Client *gate.APIClient
	Config *config.Config
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
		Client: gate.NewAPIClient(conf),
		Config: cfg,
	}

	return c
}

func (c *Client) Authenticate(ctx context.Context) (context.Context, error) {
	confAuth := c.Config.Auth

	if confAuth != nil && confAuth.Enable {
		var tok *oauth2.Token
		var err error
		if tok = confAuth.OAuth2Config.Token; tok != nil {
			conf := &oauth2.Config{
				ClientID:     confAuth.OAuth2Config.ClientID,
				ClientSecret: confAuth.OAuth2Config.ClientSecret,
				Scopes:       confAuth.OAuth2Config.Scopes,
				Endpoint: oauth2.Endpoint{
					AuthURL:  confAuth.OAuth2Config.AuthURL,
					TokenURL: confAuth.OAuth2Config.TokenURL,
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
		confAuth.OAuth2Config.Token = tok
		if err := c.Config.Write(); err != nil {
			return nil, err
		}
	}

	return ctx, nil
}
