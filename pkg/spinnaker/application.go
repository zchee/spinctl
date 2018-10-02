// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) GetApplication(ctx context.Context, name string, expand bool) error {
	opts := make(map[string]interface{})
	if expand {
		opts["expand"] = true
	}

	app, resp, err := c.client.ApplicationControllerApi.GetApplicationUsingGET(ctx, name, opts)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.Wrapf(err, "failed to get %s application", name)
	}
	defer resp.Body.Close()

	fmt.Printf("app: %#v\n", app)

	return nil
}

func (c *Client) ListApplications(ctx context.Context) error {
	apps, resp, err := c.client.ApplicationControllerApi.GetAllApplicationsUsingGET(ctx, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to get all applications")
	}
	defer resp.Body.Close()

	fmt.Printf("apps: %#v\n", apps)

	return nil
}
