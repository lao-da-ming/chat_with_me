package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
