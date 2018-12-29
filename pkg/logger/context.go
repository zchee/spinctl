// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"context"

	"go.uber.org/zap"
)

type (
	ctxLogger        struct{}
	ctxSugaredLogger struct{}
)

var (
	ctxLoggerKey        = &ctxLogger{}
	ctxSugaredLoggerKey = &ctxSugaredLogger{}
)

// NewContext return the context with zap SugaredLogger context value.
func NewContext(ctx context.Context, sugaredLogger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, ctxSugaredLoggerKey, sugaredLogger)
}

// NewContextLogger return the context with zap Logger context value.
func NewContextLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey, logger)
}

// NewContextSugaredLogger wrap of NewContext.
func NewContextSugaredLogger(ctx context.Context, sugaredLogger *zap.SugaredLogger) context.Context {
	return NewContext(ctx, sugaredLogger)
}

// FromContext extracts zap SuggeredLogger from context.
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l, ok := ctx.Value(ctxSugaredLoggerKey).(*zap.SugaredLogger)
	if !ok {
		return nil
	}
	return l
}

// FromContextLogger extracts zap Logger from context.
func FromContextLogger(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(ctxLoggerKey).(*zap.Logger)
	if !ok {
		return nil
	}
	return l
}

// FromContextSugaredLogger wrap of FromContext.
func FromContextSugaredLogger(ctx context.Context) *zap.SugaredLogger {
	return FromContext(ctx)
}
