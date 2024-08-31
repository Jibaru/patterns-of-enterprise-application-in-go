package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:product.db?cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	// Create products table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		price REAL,
		version INTEGER
	);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
