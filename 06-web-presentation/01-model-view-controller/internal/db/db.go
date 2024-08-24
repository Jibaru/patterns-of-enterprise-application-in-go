package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./todo_list.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	createTodosTable := `
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			description TEXT,
			completed BOOLEAN
		);`
	_, err = db.Exec(createTodosTable)
	if err != nil {
		return nil, fmt.Errorf("failed to create todos table: %v", err)
	}

	return db, nil
}
