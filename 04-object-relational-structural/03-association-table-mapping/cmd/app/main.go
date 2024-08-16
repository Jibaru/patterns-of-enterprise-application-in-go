package main

import (
	"fmt"
	"log"

	"github.com/jibaru/association-table-mapping/internal/db"
	"github.com/jibaru/association-table-mapping/internal/domain"
)

func main() {
	// Initialize the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("failed to set up database: %v", err)
	}

	// Create repositories
	authorRepo := domain.NewAuthorRepository(database)
	bookRepo := domain.NewBookRepository(database)
	authorBookRepo := domain.NewAuthorBookRepository(database)

	// Create a new author
	author := &domain.Author{Name: "George Orwell"}
	err = authorRepo.Insert(author)
	if err != nil {
		log.Fatalf("failed to insert author: %v", err)
	}
	fmt.Printf("Author created: %+v\n", author)

	// Create some books
	book1 := &domain.Book{Title: "1984"}
	book2 := &domain.Book{Title: "Animal Farm"}

	books := []*domain.Book{book1, book2}
	for _, book := range books {
		err = bookRepo.Insert(book)
		if err != nil {
			log.Fatalf("failed to insert book: %v", err)
		}
		fmt.Printf("Book created: %+v\n", book)

		// Associate the book with the author
		err = authorBookRepo.Associate(author.ID, book.ID)
		if err != nil {
			log.Fatalf("failed to associate author and book: %v", err)
		}
	}

	// Retrieve all books written by the author
	retrievedBooks, err := authorBookRepo.GetBooksByAuthor(author.ID)
	if err != nil {
		log.Fatalf("failed to retrieve books: %v", err)
	}

	fmt.Printf("Books by %s:\n", author.Name)
	for _, book := range retrievedBooks {
		fmt.Printf("- %s\n", book.Title)
	}
}
