// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"bytes"
	"encoding/json"

	"github.com/ghodss/yaml"
)

func parsePayload(payload interface{}, format string) ([]byte, error) {
	data, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	var buf []byte
	switch format {
	case "yaml":
		buf, err = yaml.JSONToYAML(data)
		if err != nil {
			return nil, err
		}
	case "raw":
		buf = data
	case "json":
		fallthrough // json is defaut
	default:
		b := new(bytes.Buffer)
		if err := json.Indent(b, data, "", "  "); err != nil {
			return nil, err
		}
		buf = b.Bytes()
	}

	return bytes.TrimSpace(buf), nil
}
