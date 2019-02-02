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

// NewZapLogger returns the new zap Sugared Logger.
func NewZapLogger(lv zap.AtomicLevel, opts ...zap.Option) *zap.SugaredLogger {
	cfg := newZapConfig()
	cfg.Level.SetLevel(lv.Level())
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
