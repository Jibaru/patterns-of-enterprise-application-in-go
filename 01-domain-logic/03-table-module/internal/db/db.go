package db

import "database/sql"

// Setup is a helper function to open and initialize the database
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "revenue_recognition.db")
	if err != nil {
		return nil, err
	}

	// Create the tables if they don't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS contracts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            product_id INTEGER NOT NULL,
            total_revenue REAL NOT NULL,
            date_signed DATE NOT NULL,
            FOREIGN KEY (product_id) REFERENCES products(id)
        );
        CREATE TABLE IF NOT EXISTS products (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            type TEXT NOT NULL
        );
        CREATE TABLE IF NOT EXISTS revenue_recognitions (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            contract_id INTEGER NOT NULL,
            amount REAL NOT NULL,
            date DATE NOT NULL,
            FOREIGN KEY (contract_id) REFERENCES contracts(id)
        );
    `)
	if err != nil {
		return nil, err
	}

	// Insert sample data
	_, err = db.Exec(`
        INSERT INTO products (name, type) VALUES ('Accounting Software', 'Software'), ('Consulting Service', 'Service');
        INSERT INTO contracts (product_id, total_revenue, date_signed) VALUES (1, 3000.00, '2024-01-01'), (2, 1500.00, '2024-02-01');
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}
