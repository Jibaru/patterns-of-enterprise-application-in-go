package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:customer.db?cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	// Create tables
	customerTable := `
	CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);`
	addressTable := `
	CREATE TABLE IF NOT EXISTS addresses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_id INTEGER,
		street TEXT,
		city TEXT,
		zip TEXT
	);`
	_, err = db.Exec(customerTable)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(addressTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}
