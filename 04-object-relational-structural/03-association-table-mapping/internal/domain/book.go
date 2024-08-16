package domain

import (
	"database/sql"
	"fmt"
)

// Book represents a book in the library.
type Book struct {
	ID    int    // Identity field for the book
	Title string // Title of the book
}

// BookRepository provides methods to interact with books in the database.
type BookRepository struct {
	DB *sql.DB
}

// NewBookRepository creates a new BookRepository.
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

// Insert adds a new book to the database.
func (repo *BookRepository) Insert(book *Book) error {
	query := `INSERT INTO books (title) VALUES (?)`
	result, err := repo.DB.Exec(query, book.Title)
	if err != nil {
		return fmt.Errorf("failed to insert book: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve book ID: %w", err)
	}

	book.ID = int(id)
	return nil
}

// GetByID retrieves a book by its ID.
func (repo *BookRepository) GetByID(id int) (*Book, error) {
	query := `SELECT id, title FROM books WHERE id = ?`
	row := repo.DB.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve book: %w", err)
	}

	return &book, nil
}
