package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Setup is a helper function to open and initialize the database
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "bank.db")
	if err != nil {
		return nil, err
	}

	// Create the accounts table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS accounts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            balance REAL NOT NULL
        )
    `)
	if err != nil {
		return nil, err
	}

	// Insert sample data
	_, err = db.Exec(`
        INSERT INTO accounts (name, balance) VALUES
        ('Alice', 1000.00),
        ('Bob', 500.00)
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}
