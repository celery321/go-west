package logger

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ log.Logger = (*ZapLogger2)(nil)

type structuredLoggerdefault struct {
	zapLogger *zap.SugaredLogger
}


type ZapLogger2 struct {
	LLog *zap.SugaredLogger

}


// New2 New logger initializes logger
func New2(inputLogConfig *Configuration)  *ZapLogger2  {
	return inputLogConfig.newZapLogger2()
}

func (logConfig *Configuration) newZapLogger2() *ZapLogger2 {
	var cores []zapcore.Core

	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//logLevel := getZapLevel(logConfig.LogLevel)
	writer := getPluginLogFilePath(logConfig.LogLocation)
	cores = append(cores, zapcore.NewCore(getEncoder(), writer, lowPriority))

	//logLevel2 := getZapLevel(logConfig.LogLevel2)
	writer2 := getPluginLogFilePath(logConfig.LogLocation2)
	cores = append(cores, zapcore.NewCore(getEncoder(), writer2, highPriority))

	combinedCore := zapcore.NewTee(cores...)

	Logg = zap.New(combinedCore,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	defer Logg.Sync()
	sugar := Logg.Sugar()

	return &ZapLogger2{
		LLog: sugar,
	}
}



// Log print the kv pairs log.
func (l *ZapLogger2) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.LLog.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.LLog.Desugar().Debug("", data...)
	case log.LevelInfo:
		l.LLog.Desugar().Info("", data...)
	case log.LevelWarn:
		l.LLog.Desugar().Warn("", data...)
	case log.LevelError:
		l.LLog.Desugar().Error("", data...)
	}
	return nil
}
