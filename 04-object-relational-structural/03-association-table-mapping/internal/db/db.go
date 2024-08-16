package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the SQLite database and creates the tables for authors, books, and the association table author_book.
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./library.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create Authors, Books, and Author_Book tables
	createTablesQuery := `
    CREATE TABLE IF NOT EXISTS authors (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );

    CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL
    );

    CREATE TABLE IF NOT EXISTS author_book (
        author_id INTEGER,
        book_id INTEGER,
        PRIMARY KEY (author_id, book_id),
        FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE,
        FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
    );`

	_, err = db.Exec(createTablesQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}
