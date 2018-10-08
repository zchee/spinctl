// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinnaker

import (
	"bytes"
	"encoding/json"
	"unsafe"

	"github.com/ghodss/yaml"
)

func parsePayload(payload interface{}, format string) (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(&payload); err != nil {
		return "", err
	}

	var out []byte
	switch format {
	case "yaml":
		data, err := yaml.JSONToYAML(buf.Bytes())
		if err != nil {
			return "", err
		}
		out = data
	case "raw":
		out = buf.Bytes()
	case "json":
		fallthrough // json is defaut
	default:
		data := buf.Bytes()
		buf.Reset() // re-use buf
		if err := json.Indent(buf, data, "", "  "); err != nil {
			return "", err
		}
		out = buf.Bytes()
	}
	out = bytes.TrimSpace(out)

	return *(*string)(unsafe.Pointer(&out)), nil
}
