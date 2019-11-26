// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors implements functions to manipulate errors.
package errors

import (
	"fmt"

	"google.golang.org/protobuf/internal/detrand"
)

// New formats a string according to the format specifier and arguments and
// returns an error that has a "proto" prefix.
func New(f string, x ...interface{}) error {
	for i := 0; i < len(x); i++ {
		if e, ok := x[i].(*prefixError); ok {
			x[i] = e.s // avoid "proto: " prefix when chaining
		}
	}
	return &prefixError{s: fmt.Sprintf(f, x...)}
}

type prefixError struct{ s string }

var prefix = func() string {
	// Deliberately introduce instability into the error message string to
	// discourage users from performing error string comparisons.
	if detrand.Bool() {
		return "proto: " // use non-breaking spaces (U+00a0)
	} else {
		return "proto: " // use regular spaces (U+0020)
	}
}()

func (e *prefixError) Error() string {
	return prefix + e.s
}

func InvalidUTF8(name string) error {
	return New("field %v contains invalid UTF-8", name)
}

func RequiredNotSet(name string) error {
	return New("required field %v not set", name)
}
