package router

import (
	"ginf/app/bootstrap/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(app *app.App) *gin.Engine {
	// 创建一个 Gin 实例
	router := app.Engine

	// 注册路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// 返回注册好路由的 Gin 实例
	return router
}
