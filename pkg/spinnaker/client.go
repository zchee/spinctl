// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"bytes"
	"context"
	"net/http"
	"net/http/cookiejar"
	"os/exec"

	"github.com/pkg/errors"
	"golang.org/x/net/publicsuffix"
	"golang.org/x/oauth2"
	oauth2_google "golang.org/x/oauth2/google"
	oauth2_v2 "google.golang.org/api/oauth2/v2"

	"github.com/zchee/spinctl/api/gate"
	"github.com/zchee/spinctl/pkg/auth"
	"github.com/zchee/spinctl/pkg/config"
	"github.com/zchee/spinctl/pkg/logger"
	"github.com/zchee/spinctl/pkg/version"
)

var (
	userAgent = "spinctl/" + version.Version
)

// Client represents a wrap of gate REST API client.
type Client struct {
	Client     *gate.APIClient
	Config     *config.Config
	httpClient *http.Client
}

var defaultGateConfiguration = &gate.Configuration{
	BasePath:      "http://localhost:8084",
	DefaultHeader: make(map[string]string),
	UserAgent:     userAgent,
	HTTPClient:    http.DefaultClient,
}

// Option represents a functional option pattern for gate.Configuration.
type Option func(*gate.Configuration)

// WithBasePath sets custom URL base path to gate.APIClient.
func WithBasePath(p string) Option {
	return func(c *gate.Configuration) {
		c.BasePath = p
	}
}

// WithHost sets custom Host URL to gate.APIClient.
func WithHost(host string) Option {
	return func(c *gate.Configuration) {
		c.Host = host
	}
}

// WithtHeader sets custom header to gate.APIClient DefaultHeader.
func WithtHeader(hdr map[string]string) Option {
	return func(c *gate.Configuration) {
		c.DefaultHeader = hdr
	}
}

// WithUserAgent sets custom user agent to gate.APIClient
func WithUserAgent(ua string) Option {
	return func(c *gate.Configuration) {
		c.UserAgent = ua
	}
}

// WithHTTPClient sets custom http.Client to gate.APIClient.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *gate.Configuration) {
		c.HTTPClient = hc
	}
}

// NewClient returns the new Client.
func NewClient(cfg *config.Config, opts ...Option) *Client {
	conf := defaultGateConfiguration
	for _, o := range opts {
		o(conf)
	}
	// cookiejar.New return tuple secound variable always nil
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	conf.HTTPClient.Jar = cookieJar

	if ep := cfg.Gate.Endpoint; ep != "" {
		conf.BasePath = ep
	}

	return &Client{
		Client:     gate.NewAPIClient(conf),
		Config:     cfg,
		httpClient: conf.HTTPClient,
	}
}

// Authenticate authorizes the any auth method for access to Gate REST API.
func (c *Client) Authenticate(ctx context.Context) (context.Context, error) {
	confAuth := c.Config.Auth

	if confAuth != nil && confAuth.Enable {
		switch { // nolint:gocritic
		case confAuth.OAuth2Config != nil:
			oauth2Conf := confAuth.OAuth2Config
			var tok *oauth2.Token
			var err error

			switch {
			case oauth2Conf.UseGcloud:
				cmd := exec.Command("gcloud", "auth", "application-default", "print-access-token")
				out, err := cmd.Output()
				if err != nil {
					return nil, err
				}
				out = bytes.TrimSpace(out)
				tok = &oauth2.Token{
					AccessToken:  string(out),
					RefreshToken: string(out),
				}

				conf := &oauth2.Config{
					Scopes:   []string{"openid", oauth2_v2.UserinfoEmailScope, oauth2_v2.UserinfoProfileScope},
					Endpoint: oauth2_google.Endpoint,
				}
				tokSrc := conf.TokenSource(ctx, tok)
				tok, err = tokSrc.Token()
				if err != nil {
					return nil, err
				}
			case oauth2Conf.Token != nil:
				tok = oauth2Conf.Token
				conf := &oauth2.Config{
					Scopes: oauth2Conf.Scopes,
					Endpoint: oauth2.Endpoint{
						AuthURL:  oauth2Conf.AuthURL,
						TokenURL: oauth2Conf.TokenURL,
					},
				}
				tokSrc := conf.TokenSource(ctx, tok)
				tok, err = tokSrc.Token()
				if err != nil {
					return nil, err
				}
			default:
				tok, err = auth.AuthenticateOAuth2(ctx, oauth2Conf)
				if err != nil {
					return nil, err
				}
				if !tok.Valid() {
					return nil, errors.Wrapf(err, "token is invalid: %v", tok)
				}
			}
			logger.FromContext(ctx).Debugf("Authenticate: %#v", tok)

			if err := auth.Login(c.httpClient, c.Config.Gate.Endpoint, tok); err != nil {
				return nil, err
			}
			ctx = context.WithValue(ctx, gate.ContextAccessToken, tok.AccessToken)
			confAuth.OAuth2Config.Token = tok
			if err := c.Config.Write(); err != nil {
				return nil, err
			}
		}
	}

	return ctx, nil
}
