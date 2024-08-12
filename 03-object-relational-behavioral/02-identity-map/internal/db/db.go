package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the database connection and returns a *sql.DB instance
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./drones.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Ensure the table exists
	query := `
    CREATE TABLE IF NOT EXISTS drones (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        status TEXT
    );`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return db, nil
}
