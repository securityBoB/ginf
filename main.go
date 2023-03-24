package main

import (
	"ginf/app/bootstrap/app"
	"ginf/app/bootstrap/config"
	"ginf/router"
	"log"
)

func main() {

	config.LoadAppConfig()
	App := app.NewApp()
	// 加载路由
	router.InitRouter(App)
	if err := App.Run(config.GetAppConfig().Port); err != nil {
		log.Fatalf("Failed to start the application: %v", err)
	}

}
