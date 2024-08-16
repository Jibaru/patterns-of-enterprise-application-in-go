package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the SQLite database and creates the Employments table.
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./employment.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create the Employments table
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS employments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        personId INTEGER NOT NULL,
        personName VARCHAR NOT NULL,
        start DATETIME NOT NULL,
        end DATETIME NOT NULL,
        salaryAmount DECIMAL NOT NULL,
        salaryCurrency CHAR NOT NULL
    );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return db, nil
}
