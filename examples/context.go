package main

import (
	"context"

	"github.com/cguertin14/logger"
)

func main() {
	ctx := context.Background()

	ctxLogger := logger.Initialize(logger.LoggerConfig{
		Level: "info",
	})

	// Put logger instance in context
	ctx = context.WithValue(ctx, logger.LoggerCtxKey, ctxLogger)

	foundLogger := logger.NewFromContextOrDefault(ctx)
	foundLogger.Info("hiiiiiiiiiii")
	foundLogger.Info("i existttttt")
}
