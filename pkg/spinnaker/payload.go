// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"bytes"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/jsonpath"
	"sigs.k8s.io/yaml"

	"github.com/zchee/spinctl/pkg/internal/unsafeutil"
)

func parsePayload(payload interface{}, outputFormat string) (s string, err error) {
	buf := new(bytes.Buffer)
	enc := jsoniter.ConfigFastest.NewEncoder(buf)
	enc.SetIndent("", "  ")

	if strings.HasPrefix(outputFormat, "jsonpath=") {
		tok := strings.Split(outputFormat, "=")
		outputFormat = tok[0]
		template := tok[1]

		payload, err = parseJSONPath(&payload, template)
		if err != nil {
			return "", errors.Wrap(err, "parsePayload: failed to convert payload to JSONToYAML")
		}
	}

	if err := enc.Encode(&payload); err != nil {
		return "", errors.Wrap(err, "parsePayload: failed to encoding payload")
	}

	out := buf.Bytes()
	switch outputFormat {
	case "yaml":
		out, err = yaml.JSONToYAML(buf.Bytes())
		if err != nil {
			return "", errors.Wrap(err, "parsePayload: failed to convert json payload to yaml")
		}
	case "json", "jsonpath":
		// nothing to do
	default:
		return "", errors.New("invalid output option")
	}
	out = bytes.TrimSpace(out)

	s = unsafeutil.UnsafeString(out)
	return s, nil
}

func parseJSONPath(input interface{}, template string) (interface{}, error) {
	j := jsonpath.New("json-path")
	if err := j.Parse(template); err != nil {
		return nil, errors.Errorf("failed to parsing json: %v", err)
	}

	values, err := j.FindResults(input)
	if err != nil {
		return nil, errors.Errorf("failed to parsing value from input %v using template %s: %v ", input, template, err)
	}

	if values != nil && len(values) > 0 && len(values[0]) > 0 {
		return values[0][0].Interface(), nil
	}

	return nil, errors.Errorf("failed to parsing value from input %v using template %s: %v ", input, template, err)
}
