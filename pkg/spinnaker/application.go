// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

func (c *Client) GetApplication(ctx context.Context, out io.Writer, name string, expand bool, format string) error {
	opts := make(map[string]interface{})
	if expand {
		opts["expand"] = true
	}

	payload, resp, err := c.client.ApplicationControllerApi.GetApplicationUsingGET(ctx, name, opts)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.Wrapf(err, "failed to get %s application", name)
	}
	defer resp.Body.Close()

	data, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	var buf []byte
	switch format {
	case "json":
		b := new(bytes.Buffer)
		if err := json.Indent(b, data, "", "  "); err != nil {
			return err
		}
		buf = b.Bytes()
	case "yaml":
		buf, err = yaml.JSONToYAML(data)
		if err != nil {
			return err
		}
	default:
		buf = data
	}

	fmt.Fprintln(out, string(bytes.TrimSpace(buf)))

	return nil
}

func (c *Client) ListApplications(ctx context.Context, out io.Writer, format string) error {
	payload, resp, err := c.client.ApplicationControllerApi.GetAllApplicationsUsingGET(ctx, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "failed to get all applications")
	}
	defer resp.Body.Close()

	data, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	var buf []byte
	switch format {
	case "json":
		b := new(bytes.Buffer)
		if err := json.Indent(b, data, "", "  "); err != nil {
			return err
		}
		buf = b.Bytes()
	case "yaml":
		buf, err = yaml.JSONToYAML(data)
		if err != nil {
			return err
		}
	default:
		buf = data
	}

	fmt.Fprintln(out, string(bytes.TrimSpace(buf)))

	return nil
}
