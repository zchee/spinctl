/*
 * Spinnaker API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gate

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

var ErrInvalidNullable = errors.New("nullable cannot have non-zero Value and ExplicitNull simultaneously")

// PtrBool is a helper routine that returns a pointer to given integer value.
func PtrBool(v bool) *bool { return &v }

// PtrInt is a helper routine that returns a pointer to given integer value.
func PtrInt(v int) *int { return &v }

// PtrInt32 is a helper routine that returns a pointer to given integer value.
func PtrInt32(v int32) *int32 { return &v }

// PtrInt64 is a helper routine that returns a pointer to given integer value.
func PtrInt64(v int64) *int64 { return &v }

// PtrFloat32 is a helper routine that returns a pointer to given float value.
func PtrFloat32(v float32) *float32 { return &v }

// PtrFloat64 is a helper routine that returns a pointer to given float value.
func PtrFloat64(v float64) *float64 { return &v }

// PtrString is a helper routine that returns a pointer to given string value.
func PtrString(v string) *string { return &v }

// PtrTime is helper routine that returns a pointer to given Time value.
func PtrTime(v time.Time) *time.Time { return &v }

type NullableBool struct {
	Value        bool
	ExplicitNull bool
}

func (v NullableBool) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull && v.Value:
		return nil, ErrInvalidNullable
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBool) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

type NullableInt struct {
	Value        int
	ExplicitNull bool
}

func (v NullableInt) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull && v.Value != 0:
		return nil, ErrInvalidNullable
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInt) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

type NullableInt32 struct {
	Value        int32
	ExplicitNull bool
}

func (v NullableInt32) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull && v.Value != 0:
		return nil, ErrInvalidNullable
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInt32) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

type NullableInt64 struct {
	Value        int64
	ExplicitNull bool
}

func (v NullableInt64) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull && v.Value != 0:
		return nil, ErrInvalidNullable
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInt64) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

type NullableFloat32 struct {
	Value        float32
	ExplicitNull bool
}

func (v NullableFloat32) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull && v.Value != 0.0:
		return nil, ErrInvalidNullable
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFloat32) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

type NullableFloat64 struct {
	Value        float64
	ExplicitNull bool
}

func (v NullableFloat64) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull && v.Value != 0.0:
		return nil, ErrInvalidNullable
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFloat64) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

type NullableString struct {
	Value        string
	ExplicitNull bool
}

func (v NullableString) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull && v.Value != "":
		return nil, ErrInvalidNullable
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableString) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

type NullableTime struct {
	Value        time.Time
	ExplicitNull bool
}

func (v NullableTime) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull && !v.Value.IsZero():
		return nil, ErrInvalidNullable
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return v.Value.MarshalJSON()
	}
}

func (v *NullableTime) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
