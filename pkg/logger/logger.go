// Copyright 2018 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"

	"github.com/zchee/spinctl/pkg/internal/pretty"
)

const (
	envZapDebug = "SPINCTL_ZAP_DEBUG"
	envLogLevel = "SPINCTL_LOG_LEVEL"
)

func init() {
	if err := zap.RegisterEncoder("debug", func(encoderConfig zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return NewConsoleEncoder(encoderConfig), nil
	}); err != nil {
		panic(fmt.Errorf("logger.init: %+v", err))
	}
}

func newZapConfig(opts ...zap.Option) zap.Config {
	var cfg zap.Config
	if _, ok := os.LookupEnv(envZapDebug); !ok {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.Encoding = "debug" // already registered init function
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		opts = append(opts, zap.AddCaller())
	}

	if envlv := os.Getenv(envLogLevel); envlv != "" {
		var elv zapcore.Level
		if err := elv.UnmarshalText([]byte(envlv)); err == nil {
			cfg.Level.SetLevel(elv)
		}
	}

	return cfg
}

func NewZapLogger(lv zapcore.Level, opts ...zap.Option) *zap.Logger {
	cfg := newZapConfig(opts...)
	if !cfg.Level.Enabled(lv) {
		cfg.Level.SetLevel(lv)
	}

	l, err := cfg.Build(opts...)
	if err != nil {
		panic(fmt.Errorf("logger.NewZapLogger: %+v", err))
	}

	return l
}

func NewZapSugaredLogger(lv zapcore.Level, out zapcore.WriteSyncer, opts ...zap.Option) *zap.SugaredLogger {
	cfg := newZapConfig(opts...)
	if !cfg.Level.Enabled(lv) {
		cfg.Level.SetLevel(lv)
	}
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

func RedirectStdLog(logger *zap.Logger) func() {
	return zap.RedirectStdLog(logger)
}

type consoleEncoder struct {
	zapcore.Encoder
	consoleEncoder zapcore.Encoder
}

func NewConsoleEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	color.NoColor = false // Force enabled

	cfg.StacktraceKey = ""
	cfg2 := cfg
	cfg2.NameKey = ""
	cfg2.MessageKey = ""
	cfg2.LevelKey = ""
	cfg2.CallerKey = ""
	cfg2.StacktraceKey = ""
	cfg2.TimeKey = ""
	return consoleEncoder{
		consoleEncoder: zapcore.NewConsoleEncoder(cfg),
		Encoder:        zapcore.NewJSONEncoder(cfg2),
	}
}

func (c consoleEncoder) Clone() zapcore.Encoder {
	return consoleEncoder{
		consoleEncoder: c.consoleEncoder.Clone(),
		Encoder:        c.Encoder.Clone(),
	}
}

func (c consoleEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	line, err := c.consoleEncoder.EncodeEntry(ent, nil)
	if err != nil {
		return nil, err
	}

	line2, err := c.Encoder.EncodeEntry(ent, fields)
	if err != nil {
		return nil, err
	}

	s, err := pretty.Format(line2.Bytes())
	if err != nil {
		line.AppendString("errrr")
	}

	line2.Reset()
	line2.AppendString(string(s))

	if ent.Stack != "" {
		line2.AppendByte('\n')
		line2.AppendString("Caller StackTrace\n")
		line2.AppendString(ent.Stack)
	}

	for _, field := range fields {
		switch field.Key {
		case "stacktrace":
			line2.AppendByte('\n')
			line2.AppendString("Error StackTrace\n")
			line2.AppendString(fmt.Sprintf("%v\n", field.String))
		}
	}

	parts := strings.Split(line2.String(), "\n")
	for i := range parts {
		line.AppendString("| ")
		line.AppendString(parts[i])
		line.AppendByte('\n')
	}

	return line, nil
}
