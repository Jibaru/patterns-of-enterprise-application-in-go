package scripts

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/unit-of-work/internal/db"
	"github.com/jibaru/unit-of-work/internal/domain"
)

// DeleteBooks is a Transaction Script that deletes multiple books in a single transaction
func DeleteBooks(dbConn *sql.DB, ids []int) error {
	uow := db.NewUnitOfWork(dbConn)

	// Starting the transaction
	for _, id := range ids {
		uow.RegisterDeleted(&domain.Book{ID: id})
	}

	// Commit the transaction
	err := uow.Commit()
	if err != nil {
		fmt.Printf("Failed to commit during deletion: %v\n", err)
		return uow.Rollback()
	}

	return nil
}
