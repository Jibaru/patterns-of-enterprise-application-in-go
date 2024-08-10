package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Setup is a helper function to open and initialize the database
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		return nil, err
	}

	// Create the tables if they don't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS customers (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL
        );
        CREATE TABLE IF NOT EXISTS products (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            price REAL NOT NULL
        );
        CREATE TABLE IF NOT EXISTS orders (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            customer_id INTEGER NOT NULL,
            total REAL NOT NULL,
            FOREIGN KEY (customer_id) REFERENCES customers(id)
        );
        CREATE TABLE IF NOT EXISTS order_items (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            order_id INTEGER NOT NULL,
            product_id INTEGER NOT NULL,
            quantity INTEGER NOT NULL,
            price REAL NOT NULL,
            FOREIGN KEY (order_id) REFERENCES orders(id),
            FOREIGN KEY (product_id) REFERENCES products(id)
        );
    `)
	if err != nil {
		return nil, err
	}

	// Insert sample data
	_, err = db.Exec(`
        INSERT INTO customers (name) VALUES ('John Doe');
        INSERT INTO products (name, price) VALUES ('Laptop', 1200.00), ('Mouse', 25.00);
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}
