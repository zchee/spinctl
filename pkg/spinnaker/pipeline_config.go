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

// GetPipelineConfig retrieve a pipeline configuration.
func (c *Client) GetPipelineConfig(ctx context.Context, application, pipelineName, format string) (string, error) {
	payload, resp, err := c.Client.ApplicationControllerApi.GetPipelineConfigUsingGET(ctx, application, pipelineName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get pipeline config from %s", application)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Wrapf(err, "failed to get pipeline config. status_code: %d\n", resp.StatusCode)
	}

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}

// GetPipelineConfigFromApplication retrieve a list of an application's pipeline configurations.
func (c *Client) GetPipelineConfigFromApplication(ctx context.Context, application, format string) (string, error) {
	payload, resp, err := c.Client.ApplicationControllerApi.GetPipelineConfigsForApplicationUsingGET(ctx, application)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get pipeline config from %s", application)
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

// ConvertPipelineConfigToPipelineTemplate converts a pipeline config to a pipeline template.
func (c *Client) ConvertPipelineConfigToPipelineTemplate(ctx context.Context, id string) (string, error) {
	s, resp, err := c.Client.PipelineConfigControllerApi.ConvertPipelineConfigToPipelineTemplateUsingGET(ctx, id)
	if err != nil {
		return "", errors.Wrap(err, "failed to convert pipeline config to pipeline template")
	}
	defer resp.Body.Close()

	return s, nil
}
