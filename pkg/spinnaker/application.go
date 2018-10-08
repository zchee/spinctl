// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) GetApplication(ctx context.Context, out io.Writer, name string, expand bool, format string) error {
	opts := make(map[string]interface{})
	if expand {
		opts["expand"] = true
	}

	payload, resp, err := c.client.ApplicationControllerApi.GetApplicationUsingGET(ctx, name, opts)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.Wrapf(err, "failed to get %s application", name)
	}
	defer resp.Body.Close()

	s, err := parsePayload(&payload, format)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, s)

	return nil
}

func (c *Client) ListApplications(ctx context.Context, out io.Writer, format string) error {
	payload, resp, err := c.client.ApplicationControllerApi.GetAllApplicationsUsingGET(ctx, nil)
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
