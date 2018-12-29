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

// GetVersion fetch Gate's current version.
func (c *Client) GetVersion(ctx context.Context, out io.Writer, format string) error {
	payload, resp, err := c.Client.VersionControllerApi.GetVersionUsingGET(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get version")
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// nothing to do
	default:
		return errors.Wrapf(err, "encountered an error getting version, status code: %d", resp.StatusCode)
	}

	s, err := parsePayload(&payload, format)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, s)

	return nil
}
