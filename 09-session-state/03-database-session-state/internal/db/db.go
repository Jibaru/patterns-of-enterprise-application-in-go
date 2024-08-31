package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:session.db?cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS sessions (
        session_id TEXT PRIMARY KEY,
        data TEXT
    );
    `
	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, err
	}

	return db, nil
}
