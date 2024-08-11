package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the database connection and creates necessary tables
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "employees.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = createTables(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

// createTables creates the employees table if it doesn't exist
func createTables(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS employees (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            first_name TEXT NOT NULL,
            last_name TEXT NOT NULL,
            department TEXT NOT NULL
        );
    `)
	return err
}
