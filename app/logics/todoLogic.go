package logics

import (
	"fmt"
	"myapp/models"
	"myapp/repositories"
	"myapp/services"
	"net/http"
)

// todoのインターフェース
type TodoLogicer interface {
	GetTodos(r *http.Request) ([]models.Todo, error)
}

// todoの構造体
type TodoLogic struct {
	todoRepo repositories.TodoRepository
}

// todoインターフェースを満たすtodo構造体を返す
func NewTodoLogicer() TodoLogicer {
	var tr repositories.TodoRepository

	return &TodoLogic{tr}
}

////////// todoインターフェースを満たすtodo構造体のメソッド

func (tl *TodoLogic) GetTodos(r *http.Request) (todos []models.Todo, err error) {
	// ユーザIDを取得
	userId, err := services.GetUserIdByRequestToken(r)
	if err != nil {
		fmt.Println(err)

		return todos, err
	}

	// todoを取得
	err = tl.todoRepo.GetTodos(&todos, userId)

	// エラーハンドル
	if err != nil {
		fmt.Println(err)

		return todos, err
	}

	// todo を返す
	return todos, nil
}
