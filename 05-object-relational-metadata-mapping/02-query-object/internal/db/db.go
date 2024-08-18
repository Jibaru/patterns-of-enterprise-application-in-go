package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the SQLite database and creates the products table
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./query_object.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	drop := `DROP TABLE IF EXISTS products;`
	_, err = db.Exec(drop)
	if err != nil {
		return nil, fmt.Errorf("failed to drop products table: %v", err)
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

	products := `
		INSERT INTO products (
			id,
			name,
			price
		) VALUES (1, 'Laptop', 1500), (2, 'Mouse', 15.5), (3, 'Keyboard', 40);`
	_, err = db.Exec(products)
	if err != nil {
		return nil, fmt.Errorf("failed to insert products table: %v", err)
	}

	return db, nil
}
