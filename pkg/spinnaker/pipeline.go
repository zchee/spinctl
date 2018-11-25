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

func (c *Client) GetPipeline(ctx context.Context, out io.Writer, application, pipelineName string, format string) error {
	payload, resp, err := c.Client.ApplicationControllerApi.GetPipelineConfigUsingGET(ctx, application, pipelineName)
	if err != nil {
		return errors.Wrap(err, "failed to get all applications")
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// nothing to do
	default:
		return errors.Wrapf(err, "encountered an error getting pipeline in pipeline %s with name %s, status code: %d\n", application, pipelineName, resp.StatusCode)
	}

	s, err := parsePayload(&payload, format)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, s)

	return nil
}

func (c *Client) ListPipelines(ctx context.Context, out io.Writer, name, format string) error {
	payload, resp, err := c.Client.ApplicationControllerApi.GetPipelineConfigsForApplicationUsingGET(ctx, name)
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
