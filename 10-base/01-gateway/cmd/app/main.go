package main

import (
	"fmt"
	"log"

	"github.com/jibaru/gateway/internal/gateways"
)

func main() {
	// Initialize the UserGateway with the base URL of the external API.
	userGateway := gateways.NewUserGateway("https://jsonplaceholder.typicode.com")

	// Fetch a user by their ID (e.g., user with ID 1).
	userID := 1
	user, err := userGateway.GetUserByID(userID)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	// Display the user information.
	fmt.Printf("User ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)
}
