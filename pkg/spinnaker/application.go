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

func (c *Client) ListApplications(ctx context.Context) error {
	apps, resp, err := c.client.ApplicationControllerApi.GetAllApplicationsUsingGET(ctx, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to get all applications")
	}
	defer resp.Body.Close()

	fmt.Printf("apps: %#v\n", apps)

	return nil
}
