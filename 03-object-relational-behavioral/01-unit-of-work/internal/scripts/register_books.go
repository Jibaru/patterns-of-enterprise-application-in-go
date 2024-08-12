package scripts

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/unit-of-work/internal/db"
	"github.com/jibaru/unit-of-work/internal/domain"
)

// RegisterBooks is a Transaction Script that registers multiple books in a single transaction
func RegisterBooks(dbConn *sql.DB, booksToRegister []*domain.Book) error {
	uow := db.NewUnitOfWork(dbConn)

	// Start transaction for registering books
	for _, book := range booksToRegister {
		uow.RegisterNew(book)
	}

	// Commit the transaction
	err := uow.Commit()
	if err != nil {
		fmt.Printf("Failed to commit during registration: %v\n", err)
		return uow.Rollback()
	}

	return nil
}
