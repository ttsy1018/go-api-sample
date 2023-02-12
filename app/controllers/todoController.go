package controllers

import (
	"encoding/json"
	"fmt"
	"myapp/logics"
	"myapp/services"
	"net/http"
)

type TodoController struct{}

func (tc *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	// インターフェースを取得
	todoLogic := logics.NewTodoLogicer()

	// todoを取得
	todos, err := todoLogic.GetTodos()

	// エラーハンドル
	if err != nil {
		fmt.Println(err)
		services.SendResponse(w, services.CreateErrorStringResponse("todoリストの取得に失敗しました。"), http.StatusInternalServerError)
	}

	// json encode
	responseBody, _ := json.Marshal(todos)

	// APIレスポンス
	services.SendResponse(w, responseBody, http.StatusOK)
}
