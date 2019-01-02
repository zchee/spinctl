// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"bytes"

	"github.com/ghodss/yaml"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"github.com/zchee/spinctl/pkg/unsafeutil"
)

func parsePayload(payload interface{}, format string) (string, error) {
	buf := new(bytes.Buffer)
	enc := jsoniter.ConfigFastest.NewEncoder(buf)

	// format raw prints payload directly
	if format != "raw" {
		enc.SetIndent("", "  ")
	}
	if err := enc.Encode(&payload); err != nil {
		return "", errors.Wrap(err, "parsePayload: failed to encoding payload")
	}

	out := buf.Bytes()
	switch format {
	case "yaml":
		data, err := yaml.JSONToYAML(buf.Bytes())
		if err != nil {
			return "", errors.Wrap(err, "parsePayload: failed to convert payload to JSONToYAML")
		}
		out = data
	case "raw":
		// nothing to do
	case "json":
		// nothing to do
		// format json by defaut
	}
	out = bytes.TrimSpace(out)

	return unsafeutil.UnsafeString(out), nil
}
