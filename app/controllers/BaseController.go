package controllers

import (
	"ginf/app/utils"
	"github.com/gin-gonic/gin"
)

type BaseController struct{}

// Response 封装响应
func (c *BaseController) Response(ctx *gin.Context, status int, message string, data interface{}) {
	response, _ := ctx.Get("response")
	r := response.(*utils.Response)
	r.Message = message
	r.Data = data
	r.StatusCode = status
}
