package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the SQLite database and creates the tables for Class Table Inheritance.
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./players.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create Players table (base class)
	createPlayersTableQuery := `
    CREATE TABLE IF NOT EXISTS players (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`
	_, err = db.Exec(createPlayersTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create players table: %w", err)
	}

	// Create Footballers table (subclass)
	createFootballersTableQuery := `
    CREATE TABLE IF NOT EXISTS footballers (
        id INTEGER PRIMARY KEY,
        club TEXT,
        FOREIGN KEY(id) REFERENCES players(id)
    );`
	_, err = db.Exec(createFootballersTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create footballers table: %w", err)
	}

	// Create Cricketers table (subclass)
	createCricketersTableQuery := `
    CREATE TABLE IF NOT EXISTS cricketers (
        id INTEGER PRIMARY KEY,
        batting_average REAL,
        FOREIGN KEY(id) REFERENCES players(id)
    );`
	_, err = db.Exec(createCricketersTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create cricketers table: %w", err)
	}

	// Create Bowlers table (subclass of Cricketer)
	createBowlersTableQuery := `
    CREATE TABLE IF NOT EXISTS bowlers (
        id INTEGER PRIMARY KEY,
        bowling_average REAL,
        FOREIGN KEY(id) REFERENCES cricketers(id)
    );`
	_, err = db.Exec(createBowlersTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create bowlers table: %w", err)
	}

	return db, nil
}
