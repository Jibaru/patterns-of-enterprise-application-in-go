package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var instance *sql.DB

func Setup() (*sql.DB, error) {
	var err error
	instance, err = sql.Open("sqlite3", "blog.db")
	if err != nil {
		return nil, err
	}

	// Create the posts table if it doesn't exist
	query := `
    CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL
    );`
	if _, err := instance.Exec(query); err != nil {
		return nil, err
	}

	return instance, nil
}

// Instance returns the singleton database instance
func Instance() *sql.DB {
	return instance
}
