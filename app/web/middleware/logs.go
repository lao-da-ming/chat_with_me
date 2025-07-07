package middleware

import (
	"github.com/gin-gonic/gin"
)

func LogsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
