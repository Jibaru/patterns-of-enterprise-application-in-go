package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the SQLite database and creates necessary tables
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./metadata_mapping.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	stmts := []string{
		`DROP TABLE IF EXISTS products;`,
		`DROP TABLE IF EXISTS customers;`,
	}
	for _, stmt := range stmts {
		_, err = db.Exec(stmt)
		if err != nil {
			return nil, fmt.Errorf("failed to execute: %v", err)
		}
	}

	createProductTable := `
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY,
			name TEXT,
			price REAL
		);`
	_, err = db.Exec(createProductTable)
	if err != nil {
		return nil, fmt.Errorf("failed to create products table: %v", err)
	}

	createCustomerTable := `
		CREATE TABLE IF NOT EXISTS customers (
			id INTEGER PRIMARY KEY,
			name TEXT,
			email TEXT
		);`
	_, err = db.Exec(createCustomerTable)
	if err != nil {
		return nil, fmt.Errorf("failed to create customers table: %v", err)
	}

	return db, nil
}
