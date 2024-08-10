package scripts

import (
	"database/sql"
	"fmt"
)

// Account represents a bank account
type Account struct {
	ID      int
	Name    string
	Balance float64
}

// TransferScript is responsible for handling business logic
type TransferScript struct {
	db *sql.DB
}

// NewTransfer creates a new instance of TransactionScript
func NewTransfer(db *sql.DB) *TransferScript {
	return &TransferScript{db: db}
}

// TransferMoney transfers money from one account to another
func (ts *TransferScript) TransferMoney(fromAccountID, toAccountID int, amount float64) error {
	tx, err := ts.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var fromBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromAccountID).Scan(&fromBalance)
	if err != nil {
		return fmt.Errorf("failed to get balance of account %d: %w", fromAccountID, err)
	}

	if fromBalance < amount {
		return fmt.Errorf("insufficient funds in account %d", fromAccountID)
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromAccountID)
	if err != nil {
		return fmt.Errorf("failed to debit from account %d: %w", fromAccountID, err)
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toAccountID)
	if err != nil {
		return fmt.Errorf("failed to credit to account %d: %w", toAccountID, err)
	}

	return nil
}
