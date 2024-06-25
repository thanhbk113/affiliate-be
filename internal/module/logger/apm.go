package logger

import (
	"encoding/json"

	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	LogData struct {
		Source  string
		Message string
		Data    interface{}
	}

	Map map[string]interface{}
)

var (
	zapLogger *zap.Logger
	err       error
)

// Init ...
func Init(appName, server string) {
	cfg := zap.Config{
		Encoding:      "console",
		Level:         zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:   []string{"stdout"},
		InitialFields: map[string]interface{}{"server": server, "capture": appName},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}
	zapLogger, err = cfg.Build(zap.WrapCore((&apmzap.Core{}).WrapCore))
	if err != nil {
		panic(err)
	}
}

// Debug ...
func Debug(content string, data LogData) {
	jsonData, _ := json.Marshal(data)
	zapLogger.Debug(content, zap.String("data", string(jsonData)))
}

// APM ...
func APM(fields []zap.Field, content string, data LogData) {
	jsonData, _ := json.Marshal(data)
	zapLogger.With(fields...).Info(content, zap.String("data", string(jsonData)))
}

// Error ...
func Error(content string, data LogData) {
	jsonData, _ := json.Marshal(data)
	zapLogger.Error(content, zap.String("data", string(jsonData)))
}

// Info ...
func Info(content string, data LogData) {
	jsonData, _ := json.Marshal(data)
	zapLogger.Info(content, zap.String("data", string(jsonData)))
}

// Warn ...
func Warn(content string, data LogData) {
	jsonData, _ := json.Marshal(data)
	zapLogger.Warn(content, zap.String("data", string(jsonData)))
}

// GetZapLogger ...
func GetZapLogger() *zap.Logger {
	return zapLogger
}
