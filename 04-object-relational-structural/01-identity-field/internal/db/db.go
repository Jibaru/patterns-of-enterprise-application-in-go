package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the SQLite database and creates the tables.
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./smart_devices.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create the devices table
	query := `
		CREATE TABLE IF NOT EXISTS devices (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			model TEXT,
			status TEXT
		);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create devices table: %w", err)
	}

	return db, nil
}
