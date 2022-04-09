package main

import (
	"github.com/clubcedille/logger"
)

func main() {
	logger := logger.Initialize(logger.Config{
		Level: "debug",
	})

	logger.Info("hi")
	logger.Info("heyyyyyyy")
}
