// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxLogger struct{}

var ctxLoggerKey = &ctxLogger{}

func NewContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(ctxLoggerKey).(*zap.Logger)
	if !ok {
		return nil
	}
	return l
}

func WithContext(ctx context.Context, fields ...zapcore.Field) context.Context {
	return NewContext(ctx, FromContext(ctx).With(fields...))
}
