// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (c *Client) ListApplications(ctx context.Context) error {
	apps, _, err := c.client.ApplicationControllerApi.GetAllApplicationsUsingGET(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed to get all applications")
	}
	// TODO(zchee): handle response status

	fmt.Printf("apps: %#v\n", apps)

	return nil
}
