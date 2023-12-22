package middleware

import (
	"blog-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		context.Next()
		cost := time.Since(start)
		utils.Logger.Info(context.Request.URL.Path,
			zap.Int("status", context.Writer.Status()),
			zap.String("method", context.Request.Method),
			zap.String("query", context.Request.URL.RawQuery),
			zap.String("ip", context.ClientIP()),
			zap.String("user-agent", context.Request.UserAgent()),
			zap.String("errors", context.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
