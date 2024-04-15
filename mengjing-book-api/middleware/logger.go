package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterZapLogger 日志中间件
func RegisterZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		zap.L().Info(path,
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
		)

	}
}
