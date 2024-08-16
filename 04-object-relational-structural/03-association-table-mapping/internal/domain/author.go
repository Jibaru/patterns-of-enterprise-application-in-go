package domain

import (
	"database/sql"
	"fmt"
)

// Author represents a book author.
type Author struct {
	ID   int    // Identity field for the author
	Name string // Name of the author
}

// AuthorRepository provides methods to interact with authors in the database.
type AuthorRepository struct {
	DB *sql.DB
}

// NewAuthorRepository creates a new AuthorRepository.
func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{DB: db}
}

// Insert adds a new author to the database.
func (repo *AuthorRepository) Insert(author *Author) error {
	query := `INSERT INTO authors (name) VALUES (?)`
	result, err := repo.DB.Exec(query, author.Name)
	if err != nil {
		return fmt.Errorf("failed to insert author: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve author ID: %w", err)
	}

	author.ID = int(id)
	return nil
}

// GetByID retrieves an author by their ID.
func (repo *AuthorRepository) GetByID(id int) (*Author, error) {
	query := `SELECT id, name FROM authors WHERE id = ?`
	row := repo.DB.QueryRow(query, id)

	var author Author
	err := row.Scan(&author.ID, &author.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("author with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve author: %w", err)
	}

	return &author, nil
}
