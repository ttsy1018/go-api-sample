package controllers

import (
	"fmt"
	"net/http"
)

type AppController interface {
	RootPage(w http.ResponseWriter, r *http.Request)
}

type appController struct{}

func NewAppController() AppController {
	return &appController{}
}

// api
func (apc *appController) RootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Root endpoint is hooked!")
	fmt.Fprintf(w, "Welcome to the Go Api Server")
}
