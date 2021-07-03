// package logger should be used for service logging functionality
package logger

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Log should be used for logging output, initialize it using one of the Use functions like UseZapLogger.
var Log Logger

// Logger represent common interface for logging methods
type Logger interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
}

// UseZapLogger sets Log to a Zap logger.
func UseZapLogger(logLevel string) error {
	config := fmt.Sprintf(zapConfig, logLevel)
	logger, err := zapLoggerInit(config)
	if err != nil {
		return errors.Wrap(err, "zapLoggerInit")
	}

	Log = logger
	Log.Info("Log initialized")
	return nil
}

// CloseLogger performs all necessary clean-up for Log.
func CloseLogger() {
	Log.Info("closing logger")
	switch v := Log.(type) {
	case *zap.SugaredLogger:
		zapLoggerClose(v)
	}
}
