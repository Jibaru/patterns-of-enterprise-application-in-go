package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the database and creates the necessary tables
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
            first_name TEXT NOT NULL,
            last_name TEXT NOT NULL
        );
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	return db, nil
}
