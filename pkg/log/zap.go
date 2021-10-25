package log

import (
	"fmt"
	"go-west/internal/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"strings"
)

var _  Logger = (*ZapHelper)(nil)

// ZapHelper Helper is a logger helper.
type ZapHelper struct {
	zapLogger *zap.SugaredLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.MessageKey = ""   // 使用loghelper的mesgkey
	return zapcore.NewJSONEncoder(encoderConfig)
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


// NewZapLogger NewZap NewHelper new a logger helper.
func NewZapLogger(c *conf.Server_LOG) *ZapHelper {
	var cores []zapcore.Core

	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	if c.Path == "" || c.Path == "stdout" {
		logLevel := getZapLevel("Debug")
		writer := getPluginLogFilePath(c.Path)
		cores = append(cores, zapcore.NewCore(getEncoder(), writer, logLevel))
	} else {
		accessPath := path.Join(c.Path, "access.log")
		errorPath := path.Join(c.Path, "error.log")
		accessLog := getPluginLogFilePath(accessPath)
		errorLog := getPluginLogFilePath(errorPath)
		cores = append(cores, zapcore.NewCore(getEncoder(), accessLog, lowPriority))
		cores = append(cores, zapcore.NewCore(getEncoder(), errorLog, highPriority))
	}
	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCaller(),
		zap.AddCallerSkip(4),
	)
	defer logger.Sync()
	sugar := logger.Sugar()

	return &ZapHelper{
		zapLogger: sugar,
	}
}


// Log print the kv pairs log.
func (logf *ZapHelper) Log(level Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		logf.zapLogger.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}

	switch level {
	case LevelDebug:
		logf.zapLogger.Desugar().Debug("", data...)
	case LevelInfo:
		logf.zapLogger.Desugar().Info("", data...)
	case LevelWarn:
		logf.zapLogger.Desugar().Warn("", data...)
	case LevelError:
		logf.zapLogger.Desugar().Error("", data...)
	}
	return nil
}

func (logf *ZapHelper) Debugf(format string, args ...interface{}) {
	logf.zapLogger.Debugf(format, args...)
}

func (logf *ZapHelper) Debug(format string) {
	logf.zapLogger.Desugar().Debug(format)
}

func (logf *ZapHelper) Infof(format string, args ...interface{}) {
	logf.zapLogger.Infof(format, args...)
}

func (logf *ZapHelper) Info(format string) {
	logf.zapLogger.Desugar().Info(format)
}

func (logf *ZapHelper) Warnf(format string, args ...interface{}) {
	logf.zapLogger.Warnf(format, args...)
}

func (logf *ZapHelper) Warn(format string) {
	logf.zapLogger.Desugar().Warn(format)
}

func (logf *ZapHelper) Error(format string) {
	logf.zapLogger.Desugar().Error(format)
}

func (logf *ZapHelper) Errorf(format string, args ...interface{}) {
	logf.zapLogger.Errorf(format, args...)
}

func (logf *ZapHelper) Fatalf(format string, args ...interface{}) {
	logf.zapLogger.Fatalf(format, args...)
}

func (logf *ZapHelper) Panicf(format string, args ...interface{}) {
	logf.zapLogger.Fatalf(format, args...)
}
