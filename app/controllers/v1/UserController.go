package v1

import (
	"ginf/app/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Get(ctx *gin.Context) {

	// 处理获取用户信息的逻辑
	// ...

	// 将响应数据保存到 response 实例中
	c.Response(ctx, http.StatusOK, "成功", "userdata")
}
