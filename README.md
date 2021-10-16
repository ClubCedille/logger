# logger

Simple Go logger using contexts.

## Installation

```bash
$ go get github.com/cguertin14/logger
```

## Usage

```golang
package main

import (
	"context"

	"github.com/cguertin14/logger"
)

func main() {
	ctx := context.Background()

	// Initialize logger
	ctxLogger := logger.Initialize(logger.LoggerConfig{
		Level: "info",
	})

	// Put logger instance in context
	ctx = context.WithValue(ctx, logger.LoggerCtxKey, ctxLogger)

	// Fetch logger from context
	foundLogger := logger.NewFromContextOrDefault(ctx)
	foundLogger.Info("hey there")
}

```
