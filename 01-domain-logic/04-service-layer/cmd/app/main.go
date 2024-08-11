package main

import (
	"fmt"
	"log"

	"github.com/jibaru/service-layer/internal/db"
	"github.com/jibaru/service-layer/internal/repositories"
	"github.com/jibaru/service-layer/internal/services"
)

func main() {
	// Set up the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}
	defer dbConn.Close()

	// Initialize the repository and service layers
	userRepo := repositories.NewUserRepository(dbConn)
	userService := services.NewUserService(userRepo)

	// Register a new user
	err = userService.RegisterUser("john_doe", "password123", "John", "Doe")
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}
	fmt.Println("User registered successfully!")

	// Authenticate the user
	authenticated, err := userService.AuthenticateUser("john_doe", "password123")
	if err != nil {
		log.Fatalf("Failed to authenticate user: %v", err)
	}
	if authenticated {
		fmt.Println("User authenticated successfully!")
	} else {
		fmt.Println("Invalid username or password.")
	}

	// Retrieve user details
	user, err := userService.GetUserDetails("john_doe")
	if err != nil {
		log.Fatalf("Failed to get user details: %v", err)
	}
	fmt.Printf("User details: %v\n", user)
}
