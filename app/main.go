package main

import (
	"fmt"
	"myapp/router"
	"myapp/db"
)

func main() {
	fmt.Println("run main")
	
	// DB接続
	db.Init()

	// ルータ起動
	router.Start()
}
