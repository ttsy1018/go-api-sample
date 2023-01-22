package router

import (
	"github.com/gorilla/mux"
	"myapp/controllers"
)

type AppRouter interface {
	SetAppRouting(router *mux.Router)
}

type appRouter struct {
	// AppControllerインターフフェース
	// ※AppControllerのメソッドの集合体みたいなもの
	apc controllers.AppController
}

func NewAppRouter(apc controllers.AppController) AppRouter {
	// structのアドレスを返す
	return &appRouter{apc}
}

func (apr *appRouter) SetAppRouting(router *mux.Router) {
	router.HandleFunc("/api/v1", apr.apc.RootPage).Methods("GET")
}
