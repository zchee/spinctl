// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"bytes"
	"encoding/json"
	"unsafe"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

func parsePayload(payload interface{}, format string) (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(&payload); err != nil {
		return "", errors.Wrap(err, "parsePayload: failed to encoding payload")
	}

	var out []byte
	switch format {
	case "yaml":
		data, err := yaml.JSONToYAML(buf.Bytes())
		if err != nil {
			return "", errors.Wrap(err, "parsePayload: failed to convert payload to JSONToYAML")
		}
		out = data
	case "raw":
		cb := new(bytes.Buffer)
		if err := json.Compact(cb, buf.Bytes()); err != nil {
			return "", errors.Wrap(err, "parsePayload: failed to compact for raw output")
		}
		out = cb.Bytes()
	case "json":
		fallthrough // json is defaut
	default:
		out = buf.Bytes()
	}
	out = bytes.TrimSpace(out)

	return *(*string)(unsafe.Pointer(&out)), nil
}
