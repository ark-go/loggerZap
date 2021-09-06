package main

import (
	"time"

	"go.uber.org/zap"
)

func init() {

}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "http:://ccccccc",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s--%s", "https://hhhhhh.uk", "kkkkk")
}
