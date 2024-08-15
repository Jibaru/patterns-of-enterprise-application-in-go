package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the database connection and inserts some sample data
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./ecommerce.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Ensure tables exist
	setupQueries := []string{
		`DROP TABLE customers; `,
		`DROP TABLE orders;`,
		`DROP TABLE order_lines`,
		`CREATE TABLE IF NOT EXISTS customers (id INTEGER PRIMARY KEY, name TEXT);`,
		`CREATE TABLE IF NOT EXISTS orders (id INTEGER PRIMARY KEY, customer_id INTEGER, total DECIMAL);`,
		`CREATE TABLE IF NOT EXISTS order_lines (id INTEGER PRIMARY KEY, order_id INTEGER, product_name TEXT, quantity INTEGER);`,
	}

	for _, query := range setupQueries {
		_, err = db.Exec(query)
		if err != nil {
			return nil, fmt.Errorf("failed to create tables: %w", err)
		}
	}

	// Insert sample data
	insertDataQueries := []string{
		// Insert Customers
		`INSERT INTO customers (id, name) VALUES (1, 'Alice'), (2, 'Bob');`,

		// Insert Orders
		`INSERT INTO orders (id, customer_id, total) VALUES (1, 1, 150.50), (2, 2, 99.99);`,

		// Insert Order Lines
		`INSERT INTO order_lines (id, order_id, product_name, quantity) VALUES 
            (1, 1, 'Laptop', 1),
            (2, 1, 'Mouse', 2),
            (3, 2, 'Monitor', 1),
            (4, 2, 'Keyboard', 1);`,
	}

	for _, query := range insertDataQueries {
		_, err = db.Exec(query)
		if err != nil {
			return nil, fmt.Errorf("failed to insert sample data: %w", err)
		}
	}

	return db, nil
}
