/*
 * Spinnaker API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.18.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gate

import (
	"bytes"
	"encoding/json"
)

// Version struct for Version
type Version struct {
	Version *string `json:"version,omitempty"`
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Version) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetVersionOk() (string, bool) {
	if o == nil || o.Version == nil {
		var ret string
		return ret, false
	}
	return *o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Version) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Version) SetVersion(v string) {
	o.Version = &v
}

type NullableVersion struct {
	Value        Version
	ExplicitNull bool
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
