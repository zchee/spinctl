// Copyright 2019 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// GetPipelineTemplate gets a pipeline template.
func (c *Client) GetPipelineTemplate(ctx context.Context, id, format string) (string, error) {
	payload, resp, err := c.Client.PipelineTemplatesControllerApi.GetUsingGET(ctx, id)
	if err != nil {
		return "", errors.Wrap(err, "failed to get list of pipeline template")
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// nothing to do
	default:
		return "", errors.Wrapf(err, "encountered an error getting pipeline template, id: %s, status code: %d\n", resp.StatusCode)
	}

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}
