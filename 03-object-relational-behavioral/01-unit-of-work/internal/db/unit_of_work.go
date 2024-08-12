package db

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/unit-of-work/internal/domain"
)

// UnitOfWork manages a set of operations to be committed as a single transaction
type UnitOfWork struct {
	db           *sql.DB
	newBooks     []*domain.Book
	dirtyBooks   []*domain.Book
	deletedBooks []*domain.Book
	backup       map[int]*domain.Book // Backup of books for rollback
}

// NewUnitOfWork creates a new UnitOfWork instance
func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{
		db:     db,
		backup: make(map[int]*domain.Book),
	}
}

// RegisterNew registers a new book to be inserted into the database
func (uow *UnitOfWork) RegisterNew(book *domain.Book) {
	uow.newBooks = append(uow.newBooks, book)
	fmt.Printf("Book registered as new: %v\n", *book)
}

// RegisterDirty registers a book to be updated in the database
func (uow *UnitOfWork) RegisterDirty(book *domain.Book) {
	// Backup current state before updating
	uow.backup[book.ID] = uow.getBookByID(book.ID)
	uow.dirtyBooks = append(uow.dirtyBooks, book)
	fmt.Printf("Book registered as dirty: %v\n", *book)
}

// RegisterDeleted registers a book to be deleted from the database
func (uow *UnitOfWork) RegisterDeleted(book *domain.Book) {
	// Backup current state before deleting
	uow.backup[book.ID] = uow.getBookByID(book.ID)
	uow.deletedBooks = append(uow.deletedBooks, book)
	fmt.Printf("Book registered as deleted: %v\n", *book)
}

// Commit commits all the registered operations as a single transaction
func (uow *UnitOfWork) Commit() error {
	fmt.Println("Committing transaction")
	tx, err := uow.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Insert new books
	for _, book := range uow.newBooks {
		_, err := tx.Exec("INSERT INTO books (title, author, isbn) VALUES (?, ?, ?)", book.Title, book.Author, book.ISBN)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert book: %w", err)
		}
	}

	// Update dirty books
	for _, book := range uow.dirtyBooks {
		_, err := tx.Exec("UPDATE books SET title = ?, author = ?, isbn = ? WHERE id = ?", book.Title, book.Author, book.ISBN, book.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update book: %w", err)
		}
	}

	// Delete books
	for _, book := range uow.deletedBooks {
		_, err := tx.Exec("DELETE FROM books WHERE id = ?", book.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete book: %w", err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Clear the backup after a successful commit
	uow.backup = make(map[int]*domain.Book)

	fmt.Println("Transaction committed")
	return nil
}

// Rollback rolls back the last committed operation
func (uow *UnitOfWork) Rollback() error {
	tx, err := uow.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin rollback transaction: %w", err)
	}

	// Restore deleted books
	for id, book := range uow.backup {
		if book != nil {
			_, err := tx.Exec("INSERT INTO books (id, title, author, isbn) VALUES (?, ?, ?, ?)", id, book.Title, book.Author, book.ISBN)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to restore deleted book: %w", err)
			}
		}
	}

	// Revert updates to dirty books
	for id, book := range uow.backup {
		if book != nil {
			_, err := tx.Exec("UPDATE books SET title = ?, author = ?, isbn = ? WHERE id = ?", book.Title, book.Author, book.ISBN, id)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to revert update on book: %w", err)
			}
		}
	}

	// Revert new books insertions by deleting them
	for _, book := range uow.newBooks {
		_, err := tx.Exec("DELETE FROM books WHERE isbn = ?", book.ISBN)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete new book during rollback: %w", err)
		}
	}

	// Commit the rollback transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit rollback transaction: %w", err)
	}

	// Clear the backup after a successful rollback
	uow.backup = make(map[int]*domain.Book)

	return nil
}

// getBookByID retrieves a book by ID from the database
func (uow *UnitOfWork) getBookByID(id int) *domain.Book {
	book := &domain.Book{}
	row := uow.db.QueryRow("SELECT id, title, author, isbn FROM books WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
	if err != nil {
		return nil
	}
	return book
}
