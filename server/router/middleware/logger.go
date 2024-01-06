package middleware

import (
	"blog-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		cost := time.Since(start)
		utils.Logger.Info(ctx.Request.URL.Path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
