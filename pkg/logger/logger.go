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
		panic(err)
	}
}

func NewZapLogger(l zapcore.Level, opts ...zap.Option) *zap.Logger {
	var cfg zap.Config
	if _, ok := os.LookupEnv(envZapDebug); !ok {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.Encoding = "debug" // already registered init function
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		opts = append(opts, zap.AddCaller())
	}

	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeTime = nil
	cfg.Level.SetLevel(zapcore.DPanicLevel) // not show logs normally

	if envlv := os.Getenv(envLogLevel); envlv != "" {
		var lv zapcore.Level
		if err := lv.UnmarshalText([]byte(envlv)); err == nil {
			cfg.Level.SetLevel(lv)
		}
	}

	zl, err := cfg.Build(opts...)
	if err != nil {
		panic(err)
	}

	return zl
}

func NewRedirectZapLogger(lv zapcore.Level, opts ...zap.Option) (*zap.Logger, func()) {
	log := NewZapLogger(lv)
	undo := zap.RedirectStdLog(log)

	return log, undo
}

func NewZapCore(buf *SyncBuffer) zapcore.Core {
	jsonEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "",
		CallerKey:      "", // not show
		MessageKey:     "msg",
		StacktraceKey:  "", // not show
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	})
	core := zapcore.NewCore(jsonEncoder, buf, zapcore.InfoLevel)

	return core
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
