package main

import (
	"fmt"
	"log"

	"github.com/jibaru/transaction-script/internal/db"
	"github.com/jibaru/transaction-script/internal/scripts"
)

func main() {
	// Set up the database
	db, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer db.Close()

	// Create a new instance of the TransactionScript
	ts := scripts.NewTransfer(db)

	// Perform a money transfer
	err = ts.TransferMoney(1, 2, 200.00)
	if err != nil {
		log.Fatalf("Failed to transfer money: %v", err)
	}

	fmt.Println("Money transferred successfully!")

	// Verify balances
	var balance float64
	db.QueryRow("SELECT balance FROM accounts WHERE id = 1").Scan(&balance)
	fmt.Printf("Account 1 balance: %.2f\n", balance)
	db.QueryRow("SELECT balance FROM accounts WHERE id = 2").Scan(&balance)
	fmt.Printf("Account 2 balance: %.2f\n", balance)
}
