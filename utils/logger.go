package utils

import (
	"fmt"
	"golang-rest-api/config"
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
		_config := config.GetConfig()

		var err error
		if _config.DebugMode {
			logConfig := zap.NewDevelopmentConfig()
			logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // ログレベルに色を付ける
			logger, err = logConfig.Build()
		} else {
			logger, err = zap.NewProduction()
		}

		if err != nil {
			panic("Failed to initialize zap logger: " + err.Error())
		}
	})
	return logger
}
