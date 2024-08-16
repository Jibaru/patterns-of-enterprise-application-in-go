package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the SQLite database and creates the Players table.
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./players.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create the Players table
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS players (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        club TEXT,
        batting_average REAL,
        bowling_average REAL,
        type TEXT NOT NULL
    );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return db, nil
}
