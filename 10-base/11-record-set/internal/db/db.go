package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the database and creates the necessary tables
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        DROP TABLE IF EXISTS users;
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to drop users table: %w", err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            first_name TEXT NOT NULL,
            last_name TEXT NOT NULL
        );
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	_, err = db.Exec(`
        INSERT INTO users (
            id,
            username,
            first_name,
            last_name
        ) VALUES
            ('1', 'Jose2002', 'Jose', 'Gomez'),
            ('2', 'Lua120', 'Luciana', 'Martinez'),
            ('3', 'mik', 'Michael', 'Rosas');
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to insert users table: %w", err)
	}

	return db, nil
}
