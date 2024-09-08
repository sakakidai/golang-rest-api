package utils

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func NewLogger() *zap.Logger {
	once.Do(func() {
		fmt.Println("Initialize zap logger")

		var err error
		env := os.Getenv("GO_ENV")
		if env == "development" {
			config := zap.NewDevelopmentConfig()
			config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // ログレベルに色を付ける
			logger, err = config.Build()
		} else {
			logger, err = zap.NewProduction()
		}

		if err != nil {
			panic("Failed to initialize zap logger: " + err.Error())
		}
	})
	return logger
}
