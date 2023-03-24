package router

import (
	"ginf/app/bootstrap/app"
	"ginf/app/controllers/v1"
	"ginf/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *app.App) *gin.Engine {
	// 创建一个 Gin 实例
	router := app.Engine
	router.Use(middleware.Trace())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.ResponseMiddleware())

	// v1 版本的路由
	v1Group := router.Group("/v1")
	{
		// 用户相关的路由
		userController := new(v1.UserController)
		userRouter := v1Group.Group("/user")
		{
			//userRouter.POST("/", userController.Create)
			userRouter.GET("/:id", userController.Get)
			//userRouter.PUT("/:id", userController.Update)
			//userRouter.DELETE("/:id", userController.Delete)
		}
	}
	// 返回注册好路由的 Gin 实例
	return router
}
