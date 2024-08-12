package scripts

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/unit-of-work/internal/db"
	"github.com/jibaru/unit-of-work/internal/domain"
)

// UpdateBooks is a Transaction Script that updates multiple books in a single transaction
func UpdateBooks(dbConn *sql.DB, booksToUpdate []*domain.Book) error {
	uow := db.NewUnitOfWork(dbConn)

	// Start transaction for updating books
	for _, book := range booksToUpdate {
		uow.RegisterDirty(book)
	}

	// Commit the transaction
	err := uow.Commit()
	if err != nil {
		fmt.Printf("Failed to commit during update: %v\n", err)
		return uow.Rollback()
	}

	return nil
}
