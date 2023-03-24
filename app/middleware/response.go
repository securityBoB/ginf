package middleware

import (
	"ginf/app/utils"
	"github.com/gin-gonic/gin"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := &utils.Response{
			Data: make([]interface{}, 0),
			Meta: utils.Meta{
				Trace: utils.Trace{
					TraceID:   c.MustGet("TraceID").(string),
					RequestID: c.MustGet("requestID").(string),
				},
			},
			Errors:     nil,
			StatusCode: 200,
		}
		c.Set("response", response)
		c.Next()
		c.JSON(response.StatusCode, response)
	}
}
