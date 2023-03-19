package repositories

import (
	"fmt"
	"myapp/db"
	"myapp/models"
)

type TodoRepository struct{}

// todo全件取得
func (tr *TodoRepository) GetTodos(todos *[]models.Todo, userId int) error {
	// DBからtodoを取得
	// ORMでクエリ実行 -> 指定ポインタ先の変数に格納
	result := db.DB.Where("user_id = ?", userId).Find(&todos)

	// エラーチェック
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// todo1件取得
func (tr *TodoRepository) getTodoById(todo *models.Todo, id int) error {
	fmt.Println("exec getTodoById")

	return nil
}

// todo作成
func (tr *TodoRepository) createTodo(todo *models.Todo) error {
	fmt.Println("exec createTodo")
	// todo: 作成処理

	return nil
}

// todo更新
func (tr *TodoRepository) updateTodo(todo *models.Todo) error {
	fmt.Println("exec createTodo")
	// todo: 更新処理

	return nil
}

// todo削除
func (tr *TodoRepository) deleteTodo(todo *models.Todo) error {
	fmt.Println("exec createTodo")
	// todo: 削除処理

	return nil
}
