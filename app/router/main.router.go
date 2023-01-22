package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MainRouter interface {
	StartWebServer() error
	setupRouting() *mux.Router
}

// 各routerをまとめた構造体
type mainRouter struct {
	appR AppRouter
}

func NewMainRouter(appR AppRouter) MainRouter {
	// MainRouterインターフフェース
	// ※MainRouterのメソッドの集合体みたいなもの
	return &mainRouter{appR}
}

func (mainRouter *mainRouter) setupRouting() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	mainRouter.appR.SetAppRouting(router)

	return router
}

func (mainRouter *mainRouter) StartWebServer() error {
	fmt.Println("Start Web Server")
	return http.ListenAndServe(fmt.Sprintf(":%d", 3000), mainRouter.setupRouting())
}
