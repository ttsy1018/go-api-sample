package logics

import (
	"myapp/repositories"
	"fmt"
	"myapp/models"
)

// todoのインターフェース
type TodoLogicer interface {
	GetTodos() ([]models.Todo, error)
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

func (tl *TodoLogic) GetTodos() (todos []models.Todo, err error) {
	// todoを取得
	err = tl.todoRepo.GetTodos(&todos)

	// エラーハンドル
	if err != nil {
		fmt.Println(err)

		return todos, err
	}

	// todo を返す
	return todos, nil
}