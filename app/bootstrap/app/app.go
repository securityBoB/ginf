package app

import (
	"database/sql"
	"ginf/app/bootstrap/db"
	"github.com/gin-gonic/gin"
)

type App struct {
	*gin.Engine
	DB *sql.DB
}

func NewApp() *App {
	// todo初始化配置

	// 初始化数据库连接
	mydb, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	app := &App{gin.Default(), mydb}
	return app
}

func (app *App) Run(addr string) error {
	return app.Engine.Run(addr)
}
