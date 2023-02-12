package controllers

import (
	"fmt"
	"net/http"
)

type AppController struct{}

// 疎通確認用
func (apc *AppController)RootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Root endpoint is hooked!")
	fmt.Fprintf(w, "Welcome to the Go Api Server")
}
