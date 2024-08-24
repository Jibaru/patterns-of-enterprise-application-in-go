package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the SQLite database and creates the users table
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./repository_example.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_name TEXT,
			last_name TEXT,
			email TEXT,
			age INTEGER,
			is_active BOOLEAN
		);`
	_, err = db.Exec(createUsersTable)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %v", err)
	}

	return db, nil
}
