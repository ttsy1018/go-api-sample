package main

import (
	"fmt"
	"myapp/controllers"
	"myapp/db"
	"myapp/router"
)

func main() {
	// DB接続
	db.Init()

	fmt.Println("run main")

	// controller
	appController := controllers.NewAppController()

	// router
	appRouter := router.NewAppRouter(appController)
	mainRouter := router.NewMainRouter(appRouter)

	// API起動
	mainRouter.StartWebServer()
}
