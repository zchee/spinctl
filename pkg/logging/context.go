// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logging

import (
	"context"

	"go.uber.org/zap"
)

type (
	ctxSugaredLogger struct{}
)

var (
	ctxSugaredLoggerKey = &ctxSugaredLogger{}
)

// NewContext return the context with zap SugaredLogger context value.
func NewContext(ctx context.Context, sugaredLogger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, ctxSugaredLoggerKey, sugaredLogger)
}

// FromContext extracts zap SuggeredLogger from context.
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l, ok := ctx.Value(ctxSugaredLoggerKey).(*zap.SugaredLogger)
	if !ok {
		return nil
	}
	return l
}
