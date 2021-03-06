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

// GrantedAuthority struct for GrantedAuthority
type GrantedAuthority struct {
	Authority *string `json:"authority,omitempty"`
}

// GetAuthority returns the Authority field value if set, zero value otherwise.
func (o *GrantedAuthority) GetAuthority() string {
	if o == nil || o.Authority == nil {
		var ret string
		return ret
	}
	return *o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *GrantedAuthority) GetAuthorityOk() (string, bool) {
	if o == nil || o.Authority == nil {
		var ret string
		return ret, false
	}
	return *o.Authority, true
}

// HasAuthority returns a boolean if a field has been set.
func (o *GrantedAuthority) HasAuthority() bool {
	if o != nil && o.Authority != nil {
		return true
	}

	return false
}

// SetAuthority gets a reference to the given string and assigns it to the Authority field.
func (o *GrantedAuthority) SetAuthority(v string) {
	o.Authority = &v
}

type NullableGrantedAuthority struct {
	Value        GrantedAuthority
	ExplicitNull bool
}

func (v NullableGrantedAuthority) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableGrantedAuthority) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
