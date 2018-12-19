// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/antihax/optional"
	"github.com/pkg/errors"

	"github.com/zchee/spinctl/api/gate"
)

func (c *Client) GetApplication(ctx context.Context, out io.Writer, application string, expand bool, format string) error {
	var opts gate.GetApplicationUsingGETOpts
	if expand {
		opts.Expand = optional.NewBool(true)
	}

	payload, resp, err := c.Client.ApplicationControllerApi.GetApplicationUsingGET(ctx, application, &opts)
	if err != nil {
		return errors.Wrapf(err, "failed to get %s application", application)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// nothing to do
	case http.StatusNotFound:
		return errors.Wrapf(err, "application %q not found", application)
	default:
		return errors.Wrapf(err, "encountered an error getting application, status code: %d", resp.StatusCode)
	}

	s, err := parsePayload(&payload, format)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, s)

	return nil
}

func (c *Client) ListApplications(ctx context.Context, out io.Writer, format string) error {
	payload, resp, err := c.Client.ApplicationControllerApi.GetAllApplicationsUsingGET(ctx, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to get all applications")
	}
	defer resp.Body.Close()

	s, err := parsePayload(&payload, format)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, s)

	return nil
}
