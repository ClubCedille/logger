package main

import (
	"github.com/cguertin14/logger"
)

func main() {
	logger := logger.Initialize(logger.LoggerConfig{
		Level: "debug",
	})

	logger.Info("hi")
	logger.Info("heyyyyyyy")
}
