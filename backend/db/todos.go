package db

import (
	"database/sql"
	"fmt"
	"go-todo/models"
	"time"
)

func FetchTodos(database *sql.DB) ([]models.Todo, error) {
	rows, err := database.Query("SELECT * FROM todos WHERE deleted_at IS NULL")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %w", err)
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Content, &todo.IsDone, &todo.DeletedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func CreateTodo(database *sql.DB, content string) (int64, error) {
	result, err := database.Exec("INSERT INTO todos (content) VALUES (?)", content)
	if err != nil {
		return 0, fmt.Errorf("failed to create todo: %w", err)
	}
	return result.LastInsertId()
}

func UpdateTodoContent(database *sql.DB, todoId string, content string) error {
	_, err := database.Exec("UPDATE todos SET content = ? WHERE todo_id = ?", content, todoId)
	return err
}

func DeleteTodoByID(database *sql.DB, todoId string) error {
	_, err := database.Exec("UPDATE todos SET deleted_at = ? WHERE todo_id = ?", time.Now(), todoId)
	return err
}

func MarkTodoAsDone(database *sql.DB, todoId string) error {
	_, err := database.Exec("UPDATE todos SET is_done = 1 WHERE todo_id = ?", todoId)
	return err
}
