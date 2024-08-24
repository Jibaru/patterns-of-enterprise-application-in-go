package models

import (
	"database/sql"
)

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

func GetAllTodos(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, description, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func GetTodoByID(db *sql.DB, id int) (Todo, error) {
	var todo Todo
	err := db.QueryRow("SELECT id, title, description, completed FROM todos WHERE id = ?", id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	return todo, err
}

func CreateTodo(db *sql.DB, title, description string) error {
	_, err := db.Exec("INSERT INTO todos (title, description, completed) VALUES (?, ?, ?)", title, description, false)
	return err
}

func UpdateTodo(db *sql.DB, todo Todo) error {
	_, err := db.Exec("UPDATE todos SET title = ?, description = ?, completed = ? WHERE id = ?", todo.Title, todo.Description, todo.Completed, todo.ID)
	return err
}

func DeleteTodo(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
