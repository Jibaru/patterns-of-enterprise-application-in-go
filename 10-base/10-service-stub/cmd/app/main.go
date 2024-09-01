package main

import (
	"fmt"

	"github.com/jibaru/service-stub/internal/processor"
	"github.com/jibaru/service-stub/internal/services"
)

// Main entry point of the application
func main() {
	// Create a UserServiceStub for testing
	userService := services.NewUserServiceStub()

	// Create a UserProcessor with the stubbed service
	processor := processor.NewUserProcessor(userService)

	// Process a user request
	user, err := processor.ProcessUserRequest(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Processed user: %+v\n", user)
}
