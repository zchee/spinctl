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

// GetPipeline retrieves a pipeline configuration of application.
func (c *Client) GetPipeline(ctx context.Context, application, pipelineName, format string) (string, error) {
	payload, resp, err := c.Client.ApplicationControllerApi.GetPipelineConfigUsingGET(ctx, application, pipelineName)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			return "", errors.Wrapf(err, "not found %s pipeline on %s application, status code: %d\n", pipelineName, application, resp.StatusCode)
		}
		return "", errors.Wrapf(err, "failed to get %s pipeline on %s applications", pipelineName, application)
	}
	defer resp.Body.Close()

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}

// ListPipelines retrieves a list of an application pipeline configurations.
func (c *Client) ListPipelines(ctx context.Context, name, format string) (string, error) {
	payload, resp, err := c.Client.ApplicationControllerApi.GetPipelineConfigsForApplicationUsingGET(ctx, name)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.Wrap(err, "failed to get list of pipelines")
	}
	defer resp.Body.Close()

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}

// ListPipelineConfigs gets all pipeline configs.
func (c *Client) ListPipelineConfigs(ctx context.Context, format string) (string, error) {
	payload, resp, err := c.Client.PipelineConfigControllerApi.GetAllPipelineConfigsUsingGET(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to get list pipeline configs")
	}
	defer resp.Body.Close()

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}

// GetPipelineConfigHistory gets pipeline config history.
func (c *Client) GetPipelineConfigHistory(ctx context.Context, id string, limit int32, format string) (string, error) {
	var opts gate.GetPipelineConfigHistoryUsingGETOpts
	if limit > 0 {
		opts.Limit = optional.NewInt32(limit)
	}
	payload, resp, err := c.Client.PipelineConfigControllerApi.GetPipelineConfigHistoryUsingGET(ctx, id, &opts)
	if err != nil {
		return "", errors.Wrap(err, "failed to get list pipeline configs")
	}
	defer resp.Body.Close()

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}
