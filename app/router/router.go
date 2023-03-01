package router

import (
	"fmt"
	"log"
	"myapp/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

// コントローラ構造体をフィールドに持たせる
type Controllers struct {
	apc controllers.AppController
	tc  controllers.TodoController
	ac  controllers.AuthController
}

// 実際のルーティングを定義
// ドメイン毎にファイル切り分けても良いかも
func initHandlers() {
	c := Controllers{}

	// api connect test
	router.HandleFunc("/api", c.apc.RootPage).Methods("GET")

	// auth
	router.HandleFunc("/api/auth/signin", c.ac.Signin).Methods("POST")
	router.HandleFunc("/api/auth/signup", c.ac.Signup).Methods("POST")

	// todo
	router.HandleFunc("/api/todo/list", c.tc.GetTodos).Methods("GET")
}

func Start() {
	router = mux.NewRouter()

	// ルーティングを定義
	initHandlers()

	fmt.Printf("router initialized and listening on 3000\n")

	// リクエストを待ち受ける
	log.Fatal(http.ListenAndServe(":3000", router))
}
