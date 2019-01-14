// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envZapDebug = "SPINCTL_ZAP_DEBUG"
	envLogLevel = "SPINCTL_LOG_LEVEL"
)

func newZapConfig() (cfg zap.Config) {
	if _, ok := os.LookupEnv(envZapDebug); !ok {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.DisableCaller = false
		cfg.DisableStacktrace = false
	}

	if envlv := os.Getenv(envLogLevel); envlv != "" {
		var elv zapcore.Level
		if err := elv.UnmarshalText([]byte(envlv)); err == nil {
			cfg.Level.SetLevel(elv)
		}
	}

	return cfg
}

// NewZapLogger returns the new zap Logger.
func NewZapLogger(lv zapcore.Level, opts ...zap.Option) *zap.Logger {
	cfg := newZapConfig()
	if !cfg.Level.Enabled(lv) {
		cfg.Level.SetLevel(lv)
	}

	l, err := cfg.Build(opts...)
	if err != nil {
		panic(fmt.Errorf("logger.NewZapLogger: %+v", err))
	}

	return l
}

// NewZapSugaredLogger returns the new zap Sugared Logger.
func NewZapSugaredLogger(lv zapcore.Level, opts ...zap.Option) *zap.SugaredLogger {
	cfg := newZapConfig()
	cfg.Level.SetLevel(lv)
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeTime = nil
	cfg.EncoderConfig.LevelKey = ""
	cfg.EncoderConfig.LineEnding = zapcore.DefaultLineEnding

	l, err := cfg.Build(opts...)
	if err != nil {
		panic(fmt.Errorf("logger.NewZapSugaredLogger: %+v", err))
	}

	return l.Sugar()
}

// RedirectStdLog redirects output from the standard library's package-global
// logger to the supplied logger at InfoLevel and returnn the undo function.
func RedirectStdLog(logger *zap.Logger) func() {
	return zap.RedirectStdLog(logger)
}
