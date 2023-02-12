package controllers

import (
	"encoding/json"
	"log"
	"myapp/logics"
	"myapp/services"
	"net/http"
)

type AuthController struct {
	al logics.AuthLogic
}

func (ac *AuthController) Signin(w http.ResponseWriter, r *http.Request) {
	// authenticate
	authLogic := logics.NewAuthLogicer()
	token, err := authLogic.Login(w, r)

	// erro handling
	if err != nil {
		log.Fatalln(err)

		// エラーレスポンスは返しているためreturnするだけ
		return
	}

	// response
	// json encode
	responseBody, _ := json.Marshal(token)

	// APIレスポンス
	services.SendResponse(w, responseBody, http.StatusOK)
}
