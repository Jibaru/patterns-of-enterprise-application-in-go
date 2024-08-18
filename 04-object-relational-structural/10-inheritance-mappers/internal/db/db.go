package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Importing SQLite driver
)

// Setup initializes the database and returns a connection.
func Setup() (*sql.DB, error) {
	// Connecting to an SQLite3 database
	database, err := sql.Open("sqlite3", "./players.db")
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	// Creating tables
	err = createTables(database)
	if err != nil {
		return nil, fmt.Errorf("could not create tables: %v", err)
	}

	return database, nil
}

// createTables creates all necessary tables in the database.
func createTables(db *sql.DB) error {
	stmts := []string{
		`DROP TABLE IF EXISTS players;`,
		`DROP TABLE IF EXISTS footballers;`,
		`DROP TABLE IF EXISTS cricketers;`,
		`DROP TABLE IF EXISTS bowlers;`,
		`CREATE TABLE IF NOT EXISTS players (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS footballers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			club TEXT,
			FOREIGN KEY(id) REFERENCES players(id)
		);`,
		`CREATE TABLE IF NOT EXISTS cricketers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			batting_average REAL,
			FOREIGN KEY(id) REFERENCES players(id)
		);`,
		`CREATE TABLE IF NOT EXISTS bowlers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			bowling_average REAL,
			FOREIGN KEY(id) REFERENCES cricketers(id)
		);`,
	}

	for _, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("could not execute: %v", err)
		}
	}

	return nil
}
