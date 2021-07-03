package logger

import (
	"encoding/json"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const zapConfig = `{
	"level": "%s",
	"DisableCaller": false,
	"encoding": "json",
	"outputPaths": ["stdout"],
	"errorOutputPaths": ["stderr"],
	"encoderConfig": {
		"timeKey":        	"ts",
		"levelKey":       	"level",
		"messageKey":     	"msg",
		"nameKey":        	"name",
		"stacktraceKey":  	"stacktrace",
		"callerKey":      	"caller",
		"lineEnding":     	"\n",
		"timeEncoder":     	"rfc3339",
		"levelEncoder":    	"lowercaseLevel",
		"durationEncoder": 	"stringDuration",
		"callerEncoder":   	"shortCaller"
	}
}`

// zapLoggerInit creates a zap SugaredLogger.
func zapLoggerInit(config string) (*zap.SugaredLogger, error) {
	rawJSON := []byte(config)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return nil, errors.Wrap(err, "json Unmarshal cfg")
	}

	zLogger, err := cfg.Build()
	if err != nil {
		return nil, errors.Wrap(err, "cfg Build")
	}

	sugar := zLogger.Sugar()
	sugar.Debug("zap logger initialized")
	return sugar, nil
}

// zapLoggerClose performs all necessary clean-up for a zap SugaredLogger.
func zapLoggerClose(logger *zap.SugaredLogger) {
	_ = logger.Sync()
}
