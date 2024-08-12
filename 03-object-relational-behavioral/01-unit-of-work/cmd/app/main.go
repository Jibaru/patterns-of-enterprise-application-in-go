package main

import (
	"fmt"
	"log"

	"github.com/jibaru/unit-of-work/internal/db"
	"github.com/jibaru/unit-of-work/internal/domain"
	"github.com/jibaru/unit-of-work/internal/scripts"
)

func main() {
	// Initialize the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	defer dbConn.Close()

	// Register new books
	err = scripts.RegisterBooks(dbConn, []*domain.Book{
		{Title: "The Pragmatic Programmer", Author: "Andrew Hunt", ISBN: "978-0201616224"},
		{Title: "Clean Code", Author: "Robert C. Martin", ISBN: "978-0132350884"},
		{Title: "Refactoring", Author: "Martin Fowler", ISBN: "978-0201485677"},
	})
	if err != nil {
		log.Fatalf("Failed to register books: %v", err)
	}

	// Update existing books
	err = scripts.UpdateBooks(dbConn, []*domain.Book{
		{ID: 1, Title: "The Pragmatic Programmer: 20th Anniversary Edition", Author: "Andrew Hunt", ISBN: "978-0135957059"},
		{ID: 2, Title: "Clean Code: A Handbook of Agile Software Craftsmanship", Author: "Robert C. Martin", ISBN: "978-0132350884"},
	})
	if err != nil {
		log.Fatalf("Failed to update books: %v", err)
	}

	// Delete books
	err = scripts.DeleteBooks(dbConn, []int{3, 4})
	if err != nil {
		log.Fatalf("Failed to delete books: %v", err)
	}

	fmt.Println("Transaction scripts executed successfully.")
}
