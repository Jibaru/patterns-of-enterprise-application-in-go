package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:orders.db?cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	// Create orders table
	orderTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_name TEXT,
		quantity INTEGER
	);`
	_, err = db.Exec(orderTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}
