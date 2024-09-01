package sqlite

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStorage implements the Storage interface for SQLite
type SQLiteStorage struct {
	db *sql.DB
}

// NewSQLiteStorage creates a new instance of SQLiteStorage
func NewSQLiteStorage(dbFile string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	// Ensure the table exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS data (key TEXT PRIMARY KEY, value TEXT)")
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

// Save stores the key-value pair in SQLite
func (s *SQLiteStorage) Save(key string, value string) error {
	_, err := s.db.Exec("INSERT OR REPLACE INTO data (key, value) VALUES (?, ?)", key, value)
	return err
}

// Load retrieves the value associated with the key from SQLite
func (s *SQLiteStorage) Load(key string) (string, error) {
	var value string
	err := s.db.QueryRow("SELECT value FROM data WHERE key = ?", key).Scan(&value)
	if err == sql.ErrNoRows {
		return "", errors.New("key not found")
	}
	return value, err
}
