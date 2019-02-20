// Copyright 2019 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/antihax/optional"
	"github.com/pkg/errors"

	"github.com/zchee/spinctl/api/gate"
)

// GetPipelineTemplate gets a pipeline template.
func (c *Client) GetPipelineTemplate(ctx context.Context, id, format string) (string, error) {
	payload, resp, err := c.Client.V2PipelineTemplatesControllerApi.GetUsingGET2(ctx, id)
	if err != nil {
		return "", errors.Wrap(err, "failed to get list of pipeline template")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
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
	var opts gate.ListUsingGET1Opts
	if scopes != nil {
		opts.Scopes = optional.NewInterface(scopes)
	}

	payload, resp, err := c.Client.V2PipelineTemplatesControllerApi.ListUsingGET1(ctx, &opts)
	if err != nil {
		return "", errors.Wrap(err, "failed to get list of pipeline templates")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Wrapf(err, "encountered an error getting list of pipeline templates, scopes %s, status code: %d\n", scopes, resp.StatusCode)
	}

	s, err := parsePayload(&payload, format)
	if err != nil {
		return "", err
	}

	return s, nil
}

// SavePipelineTemplates saves pipeline templates.
func (c *Client) SavePipelineTemplates(ctx context.Context, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	var jsonContent map[string]interface{}
	if err = json.NewDecoder(f).Decode(&jsonContent); err != nil {
		return err
	}

	resp, err := c.Client.PipelineTemplatesControllerApi.CreateUsingPOST(ctx, jsonContent)
	if err != nil {
		return errors.Wrap(err, "failed to save of pipeline templates")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrapf(err, "encountered an error save of pipeline templates, filepath %s, status code: %d\n", filepath, resp.StatusCode)
	}

	return nil
}
