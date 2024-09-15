package utils

import (
	"fmt"
	"golang-rest-api/config"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// グローバルなログは一回のみ生成する
func InitLogger() *zap.Logger {
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

		gin.DefaultWriter = zap.NewStdLog(logger).Writer()

		if err != nil {
			panic("Failed to initialize zap logger: " + err.Error())
		}
	})
	return logger
}

func Logger() *zap.Logger {
	return logger
}
