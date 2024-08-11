package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jibaru/table-module/internal/db"
	"github.com/jibaru/table-module/internal/modules"
)

func main() {
	// Set up the database
	db, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer db.Close()

	// Create a new instance of the ContractModule
	cm := modules.NewContract(db)

	// Recognize revenue for the first contract
	err = cm.RecognizeRevenue(1)
	if err != nil {
		log.Fatalf("Failed to recognize revenue: %v", err)
	}

	fmt.Println("Revenue recognized successfully!")

	// Retrieve recognized revenue as of today
	rm := modules.NewRevenueRecognition(db)
	total, err := rm.RecognizedRevenue(1, time.Now())
	if err != nil {
		log.Fatalf("Failed to retrieve recognized revenue: %v", err)
	}

	fmt.Printf("Total recognized revenue for contract 1: %.2f\n", total)
}
