// Copyright 2019 The spinctl Authors. All rights reserved.
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
		return "", errors.Wrapf(err, "encountered an error getting pipeline template, status code: %d\n", resp.StatusCode)
	}

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}

// ListPipelineTemplates lists pipeline templates.
func (c *Client) ListPipelineTemplates(ctx context.Context, scopes []string, format string) (string, error) {
	var opts gate.ListUsingGETOpts
	if scopes != nil {
		opts.Scopes = optional.NewInterface(scopes)
	}

	payload, resp, err := c.Client.PipelineTemplatesControllerApi.ListUsingGET(ctx, &opts)
	if err != nil {
		return "", errors.Wrap(err, "failed to get list of pipeline templates")
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// nothing to do
	default:
		return "", errors.Wrapf(err, "encountered an error getting list of pipeline templates, scopes %s, status code: %d\n", scopes, resp.StatusCode)
	}

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}
