package logger

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

var _ log.Logger = (*structuredLogger)(nil)


// New2 New logger initializes logger
func New2(inputLogConfig *Configuration)  *structuredLogger  {
	return inputLogConfig.newZapLogger()
}


// Log print the kv pairs log.
func (l *structuredLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.zapLogger.Desugar().Debug("", data...)

	case log.LevelInfo:
		l.zapLogger.Desugar().Info("", data...)
	case log.LevelWarn:
		l.zapLogger.Desugar().Warn("", data...)
	case log.LevelError:
		l.zapLogger.Desugar().Error("", data...)
	}
	return nil
}
