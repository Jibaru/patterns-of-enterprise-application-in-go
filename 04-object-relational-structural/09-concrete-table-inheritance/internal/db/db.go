package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the SQLite database and creates the tables for Concrete Table Inheritance.
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./players.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create Footballers table (with all fields including inherited ones)
	createFootballersTableQuery := `
    CREATE TABLE IF NOT EXISTS footballers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        club TEXT
    );`
	_, err = db.Exec(createFootballersTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create footballers table: %w", err)
	}

	// Create Cricketers table (with all fields including inherited ones)
	createCricketersTableQuery := `
    CREATE TABLE IF NOT EXISTS cricketers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        batting_average REAL
    );`
	_, err = db.Exec(createCricketersTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create cricketers table: %w", err)
	}

	// Create Bowlers table (with all fields including inherited ones)
	createBowlersTableQuery := `
    CREATE TABLE IF NOT EXISTS bowlers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        batting_average REAL,
        bowling_average REAL
    );`
	_, err = db.Exec(createBowlersTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create bowlers table: %w", err)
	}

	return db, nil
}
