package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		c.Request.Header.Set("X-Trace-Id", traceID)
		c.Next()
	}
}
