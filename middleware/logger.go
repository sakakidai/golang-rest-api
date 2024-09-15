package middleware

import (
	"golang-rest-api/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		// グローバルなloggerにリクエストIDを付加したloggerを作成
		logger := utils.Logger().With(zap.String("request_id", requestID))
		// loggerをコンテキストに保存
		c.Set("logger", logger)

		start := time.Now()

		c.Next()

		latency := time.Since(start)

		logger.Info("Request processed",
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status_code", c.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("error_message", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}
