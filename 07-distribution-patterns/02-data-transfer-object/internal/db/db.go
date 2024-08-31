package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:user.db?cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	// Create users table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
