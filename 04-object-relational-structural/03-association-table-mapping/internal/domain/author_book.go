package domain

import (
	"database/sql"
	"fmt"
)

// AuthorBookRepository provides methods to interact with the author_book association table.
type AuthorBookRepository struct {
	DB *sql.DB
}

// NewAuthorBookRepository creates a new AuthorBookRepository.
func NewAuthorBookRepository(db *sql.DB) *AuthorBookRepository {
	return &AuthorBookRepository{DB: db}
}

// Associate adds an entry in the author_book table linking an author and a book.
func (repo *AuthorBookRepository) Associate(authorID, bookID int) error {
	query := `INSERT INTO author_book (author_id, book_id) VALUES (?, ?)`
	_, err := repo.DB.Exec(query, authorID, bookID)
	if err != nil {
		return fmt.Errorf("failed to associate author and book: %w", err)
	}
	return nil
}

// GetBooksByAuthor retrieves all books written by a specific author.
func (repo *AuthorBookRepository) GetBooksByAuthor(authorID int) ([]*Book, error) {
	query := `
    SELECT b.id, b.title
    FROM books b
    JOIN author_book ab ON ab.book_id = b.id
    WHERE ab.author_id = ?`

	rows, err := repo.DB.Query(query, authorID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve books: %w", err)
	}
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title)
		if err != nil {
			return nil, fmt.Errorf("failed to scan book: %w", err)
		}
		books = append(books, &book)
	}

	return books, nil
}
