package main

import (
	"log"

	"github.com/jibaru/repository/internal/db"
	"github.com/jibaru/repository/internal/domain"
	"github.com/jibaru/repository/internal/repository"
)

func main() {
	// Set up the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer database.Close()

	// Create a new UserRepository
	userRepo := repository.NewUserRepository(database)

	// Create a new user
	user := &domain.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Age:       30,
		IsActive:  true,
	}

	// Save the user
	err = userRepo.Save(user)
	if err != nil {
		log.Fatalf("Failed to save user: %v", err)
	}

	// Retrieve the user by ID
	savedUser, err := userRepo.FindById(user.ID)
	if err != nil {
		log.Fatalf("Failed to find user: %v", err)
	}
	log.Printf("Retrieved User: %+v\n", savedUser)

	// Update the user
	savedUser.Age = 31
	err = userRepo.Update(savedUser)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	// Retrieve all users
	users, err := userRepo.FindAll()
	if err != nil {
		log.Fatalf("Failed to retrieve all users: %v", err)
	}
	for _, u := range users {
		log.Printf("User: %+v\n", u)
	}

	// Delete the user
	err = userRepo.Delete(savedUser.ID)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}

	log.Println("User repository operations completed successfully.")
}
