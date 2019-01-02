// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// GetPipeline retrieves a pipeline configuration of application.
func (c *Client) GetPipeline(ctx context.Context, application, pipelineName, format string) (string, error) {
	payload, resp, err := c.Client.ApplicationControllerApi.GetPipelineConfigUsingGET(ctx, application, pipelineName)
	if err != nil {
		return "", errors.Wrap(err, "failed to get all applications")
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// nothing to do
	default:
		return "", errors.Wrapf(err, "encountered an error getting pipeline in pipeline %s with name %s, status code: %d\n", application, pipelineName, resp.StatusCode)
	}

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
		return "", errors.Wrap(err, "failed to get all applications")
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
		return "", errors.Wrap(err, "failed to get all applications")
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// nothing to do
	default:
		return "", errors.Wrapf(err, "encountered an error getting list of pipeline config, status code: %d\n", resp.StatusCode)
	}

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}
