// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/pkg/errors"

	"github.com/zchee/spinctl/api/gate"
)

// GetApplication retrieves an application details.
func (c *Client) GetApplication(ctx context.Context, application string, expand bool, output string) (string, error) {
	var opts gate.GetApplicationUsingGETOpts
	if expand {
		opts.Expand = optional.NewBool(true)
	}

	payload, resp, err := c.Client.ApplicationControllerApi.GetApplicationUsingGET(ctx, application, &opts)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get %s application", application)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// nothing to do
	case http.StatusNotFound:
		return "", errors.Wrapf(err, "application %q not found", application)
	default:
		return "", errors.Wrapf(err, "encountered an error getting application, status code: %d", resp.StatusCode)
	}

	s, err := parsePayload(&payload, output)
	if err != nil {
		return "", err
	}

	return s, nil
}

// ListApplications retrieves a list of applications.
func (c *Client) ListApplications(ctx context.Context, output string) (string, error) {
	payload, resp, err := c.Client.ApplicationControllerApi.GetAllApplicationsUsingGET(ctx, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.Wrap(err, "failed to get all applications")
	}
	defer resp.Body.Close()

	s, err := parsePayload(&payload, output)
	if err != nil {
		return "", err
	}

	return s, nil
}
