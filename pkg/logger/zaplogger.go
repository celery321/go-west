// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	log2 "github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type structuredLogger struct {
	zapLogger *zap.SugaredLogger
}

type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

// getZapLevel converts log level string to zapcore.Level
func getZapLevel(inputLogLevel string) zapcore.Level {
	lvl := strings.ToLower(inputLogLevel)

	switch lvl {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func (logf *structuredLogger) Debugf(format string, args ...interface{}) {
	logf.zapLogger.Debugf(format, args...)
}

func (logf *structuredLogger) Debug(format string) {
	logf.zapLogger.Desugar().Debug(format)
}

func (logf *structuredLogger) Infof(format string, args ...interface{}) {
	logf.zapLogger.Infof(format, args...)
}

func (logf *structuredLogger) Info(format string) {
	logf.zapLogger.Desugar().Info(format)
}

func (logf *structuredLogger) Warnf(format string, args ...interface{}) {
	logf.zapLogger.Warnf(format, args...)
}

func (logf *structuredLogger) Warn(format string) {
	logf.zapLogger.Desugar().Warn(format)
}

func (logf *structuredLogger) Error(format string) {
	logf.zapLogger.Desugar().Error(format)
}

func (logf *structuredLogger) Errorf(format string, args ...interface{}) {
	logf.zapLogger.Errorf(format, args...)
}

func (logf *structuredLogger) Fatalf(format string, args ...interface{}) {
	logf.zapLogger.Fatalf(format, args...)
}

func (logf *structuredLogger) Panicf(format string, args ...interface{}) {
	logf.zapLogger.Fatalf(format, args...)
}

func (logf *structuredLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := logf.zapLogger.With(f...)
	return &structuredLogger{newLogger}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func (logConfig *Configuration) newZapLogger() *structuredLogger {
	var cores []zapcore.Core

	logLevel := getZapLevel(logConfig.LogLevel)

	writer := getPluginLogFilePath(logConfig.LogLocation)

	cores = append(cores, zapcore.NewCore(getEncoder(), writer, logLevel))

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)
	defer logger.Sync()
	sugar := logger.Sugar()

	return &structuredLogger{
		zapLogger: sugar,
	}
}

// getPluginLogFilePath returns the writer
func getPluginLogFilePath(logFilePath string) zapcore.WriteSyncer {
	var writer zapcore.WriteSyncer

	if logFilePath == "" {
		writer = zapcore.Lock(os.Stderr)
	} else if strings.ToLower(logFilePath) != "stdout" {
		writer = getLogWriter(logFilePath)
	} else {
		writer = zapcore.Lock(os.Stdout)
	}

	return writer
}

//getLogWriter is for lumberjack
func getLogWriter(logFilePath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// DefaultLogger creates and returns a new default logger.
func DefaultLogger() Logger {
	productionConfig := zap.NewProductionConfig()
	productionConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	productionConfig.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		_, caller.File, caller.Line, _ = runtime.Caller(8)
		enc.AppendString(caller.FullPath())
	}
	logger, _ := productionConfig.Build()
	defer logger.Sync()
	sugar := logger.Sugar()
	return &structuredLogger{
		zapLogger: sugar,
	}
}

// DefaultLogger2 DefaultLogger creates and returns a new default logger.
func DefaultLogger2() *ZapLogger {
	productionConfig := zap.NewProductionConfig()
	productionConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	productionConfig.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		_, caller.File, caller.Line, _ = runtime.Caller(8)
		enc.AppendString(caller.FullPath())
	}
	logger, _ := productionConfig.Build()
	defer logger.Sync()
	sugar := logger.Sugar()
	return &ZapLogger{
		log: sugar.Desugar(),
	}
}

func (l *ZapLogger) Log(level log2.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log2.LevelDebug:
		l.log.Debug("", data...)
	case log2.LevelInfo:
		l.log.Info("", data...)
	case log2.LevelWarn:
		l.log.Warn("", data...)
	case log2.LevelError:
		l.log.Error("", data...)
	}
	return nil
}
