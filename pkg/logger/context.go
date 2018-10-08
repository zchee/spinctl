// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	ctxLogger        struct{}
	ctxSugaredLogger struct{}
)

var (
	ctxLoggerKey        = &ctxLogger{}
	ctxSugaredLoggerKey = &ctxSugaredLogger{}
)

func NewContext(ctx context.Context, sugaredLogger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, ctxSugaredLoggerKey, sugaredLogger)
}

func NewContextLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey, logger)
}

func NewContextSugaredLogger(ctx context.Context, sugaredLogger *zap.SugaredLogger) context.Context {
	return NewContext(ctx, sugaredLogger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	l, ok := ctx.Value(ctxSugaredLoggerKey).(*zap.SugaredLogger)
	if !ok {
		return nil
	}
	return l
}

func FromContextLogger(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(ctxLoggerKey).(*zap.Logger)
	if !ok {
		return nil
	}
	return l
}

func FromContextSugaredLogger(ctx context.Context) *zap.SugaredLogger {
	return FromContext(ctx)
}

func WithContext(ctx context.Context, fields ...zapcore.Field) context.Context {
	return NewContextLogger(ctx, FromContextLogger(ctx).With(fields...))
}
