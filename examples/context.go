package main

import (
	"context"

	"github.com/clubcedille/logger"
)

func main() {
	ctx := context.Background()

	ctxLogger := logger.Initialize(logger.Config{
		Level: "info",
	})

	// Put logger instance in context
	ctx = context.WithValue(ctx, logger.CtxKey, ctxLogger)

	foundLogger := logger.NewFromContextOrDefault(ctx)
	foundLogger.Info("hiiiiiiiiiii")
	foundLogger.Info("i existttttt")
}
