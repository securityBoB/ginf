package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		c.Set("TraceID", traceID)
		c.Writer.Header().Set("X-Trace-ID", traceID)
		fmt.Printf("Trace ID: %s\n", traceID)
		c.Next()
	}
}
