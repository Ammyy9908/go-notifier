package logger

import (
	"context"
	"errors"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

const (
	CorrelationID            = "X-Correlation-ID"
	errCorrelationIDNotFound = "correlation id was not found in context"
)

var appName string
var appEnv string

// InitializeLogger initializes global settings for the logger, including app name and version.
func InitializeLogger(name, env string) {
	appName = name
	appEnv = env
}

// New returns a new zap.SugaredLogger instance with predefined encoding and app metadata.
func New(minLogLevel LogLevel) *zap.SugaredLogger {
	return zap.New(getCore(minLogLevel), zap.AddCaller()).With(
		zap.String("appName", appName),
		zap.String("appEnv", appEnv),
	).Sugar()
}

// NonSugaredLogger returns a new zap.Logger instance with predefined encoding and app metadata.
func NonSugaredLogger(minLogLevel LogLevel) *zap.Logger {
	return zap.New(getCore(minLogLevel), zap.AddCaller()).With(
		zap.String("appName", appName),
		zap.String("appEnv", appEnv),
	)
}

// WithCorrelation returns a zap.SugaredLogger instance with Correlation ID as a field from context and app metadata.
func WithCorrelation(ctx context.Context, minLogLevel LogLevel) (*zap.SugaredLogger, error) {
	core := getCore(minLogLevel)
	newLogger := zap.New(core, zap.AddCaller()).Sugar().With(
		zap.String("appName", appName),
		zap.String("appEnv", appEnv),
	)
	ctxCorrelationID, ok := ctx.Value(CorrelationID).(string)
	if !ok {
		return newLogger, errors.New(errCorrelationIDNotFound)
	}
	return newLogger.With(zap.String(CorrelationID, ctxCorrelationID)), nil
}

// AddCorrelation adds a correlation from context to existing logger with app metadata if correlation ID exists.
func AddCorrelation(ctx context.Context, logger *zap.SugaredLogger) *zap.SugaredLogger {
	ctxCorrelationID, ok := ctx.Value(CorrelationID).(string)
	if !ok {
		logger.Error(errCorrelationIDNotFound)
		return logger
	}
	return logger.With(zap.String(CorrelationID, ctxCorrelationID))
}

// getCore returns a zapcore.Core with a log level based on project configuration.
func getCore(logLevel LogLevel) zapcore.Core {
	return zapcore.NewCore(getEncoder(), zapcore.AddSync(os.Stdout), getLogLevel(logLevel))
}

// getEncoder returns a JSON encoder with custom configuration.
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogLevel returns zapcore.Level based on specified log level.
func getLogLevel(logLevel LogLevel) zapcore.Level {
	switch logLevel {
	case Debug:
		return zapcore.DebugLevel
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Error:
		return zapcore.ErrorLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// GetLogger returns a logger instance, with or without context-based correlation.
func GetLogger(ctx ...context.Context) *zap.SugaredLogger {
	logLevel := LogLevel(viper.GetInt("log.level"))
	if len(ctx) > 0 {
		log, err := WithCorrelation(ctx[0], logLevel)
		if err != nil {
			log.Warn("Failed to add correlation ID", zap.Error(err))
		}
		return log
	}
	return New(logLevel)
}
