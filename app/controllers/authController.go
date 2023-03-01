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

// ログイン
func (ac *AuthController) Signin(w http.ResponseWriter, r *http.Request) {
	// authenticate
	authLogic := logics.NewAuthLogicer()
	token, err := authLogic.Signin(w, r)

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

// 会員登録
func (ac *AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	// authenticate
	authLogic := logics.NewAuthLogicer()
	token, err := authLogic.Signup(w, r)

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
