package logger

import (
	"context"
	"io"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// CtxKey -
	// The context key for the Logger instance.
	CtxKey = "logger"
)

var (
	// CLIFormatter
	//
	// Default fomatter for CLIs such as tools used on local machines
	CLIFormatter = &log.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	}

	// ServiceFormatter
	//
	// Default formatter for online services such as microservices
	ServiceFormatter = &log.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}

	defaultLogger log.FieldLogger
	defaultLevel  = log.InfoLevel
	defaultOutput = os.Stdout

	// defaultFormatter is CLIFormatter
	defaultFormatter = CLIFormatter

	syncOnce = sync.Once{}
)

// Config -
// The configuration used by a Logger.
type Config struct {
	Level     string
	Output    io.Writer
	Formatter log.Formatter
}

// Initialize -
// Initializes a new Logger from a given configuration.
// If config is empty, default settings are used.
func Initialize(config Config) log.FieldLogger {
	syncOnce.Do(func() {
		newLogger := &log.Logger{}
		newLogger.SetNoLock()

		// Set Formatter
		if config.Formatter == nil {
			newLogger.SetFormatter(defaultFormatter)
		} else {
			newLogger.SetFormatter(config.Formatter)
		}

		// Set Output
		if config.Output == nil {
			newLogger.SetOutput(defaultOutput)
		} else {
			newLogger.SetOutput(config.Output)
		}

		// Set Level
		if lvl, err := log.ParseLevel(config.Level); err != nil {
			newLogger.SetLevel(defaultLevel)
		} else {
			newLogger.SetLevel(lvl)
		}

		// Only add method as field if debugging.
		if newLogger.GetLevel() == log.DebugLevel {
			// Add calling method as a field
			newLogger.SetReportCaller(true)
		}

		// Assign struct to outgoing interface
		defaultLogger = newLogger
	})
	return defaultLogger
}

// NewFromContextOrDefault -
// Returns a Logger from a given Context instance.
// If the Context doesn't contain the Logger,
// a default Logger is returned.
func NewFromContextOrDefault(ctx context.Context) log.FieldLogger {
	if ctx == nil {
		return initDefaultLogger()
	}

	// Fetch logger from context
	ctxVal := ctx.Value(CtxKey)
	if ctxVal == nil {
		return initDefaultLogger()
	}

	// Cast logger to FieldLogger interface
	if logger, ok := ctxVal.(log.FieldLogger); ok {
		return logger
	}

	return initDefaultLogger()
}

func initDefaultLogger() log.FieldLogger {
	Initialize(Config{})
	return defaultLogger
}
