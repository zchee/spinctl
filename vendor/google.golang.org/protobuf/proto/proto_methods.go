// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoreflect build tag disables use of fast-path methods.
// +build !protoreflect

package proto

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

func protoMethods(m protoreflect.Message) *protoiface.Methods {
	if x, ok := m.(protoiface.Methoder); ok {
		return x.ProtoMethods()
	}
	return nil
}
