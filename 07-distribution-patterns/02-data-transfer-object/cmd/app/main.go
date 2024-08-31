package main

import (
	"fmt"
	"log"

	"github.com/jibaru/data-transfer-object/internal/db"
	"github.com/jibaru/data-transfer-object/internal/services"
)

func main() {
	// Initialize the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Initialize the user service
	userService := services.NewUserService(dbConn)

	// Create a new user
	newUser, err := userService.CreateUser("John Doe", "john.doe@example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New User Created: %+v\n", newUser)

	// Get user details
	userDTO, err := userService.GetUser(newUser.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User Details: %+v\n", userDTO)
}
