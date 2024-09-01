package main

import (
	"fmt"
	"time"

	"github.com/jibaru/value-object/internal/domain/entities"
	"github.com/jibaru/value-object/internal/domain/values"
)

func main() {
	// Create value objects
	name, err := values.NewName("John", "Doe")
	if err != nil {
		panic(err)
	}

	email, err := values.NewEmail("john.doe@example.com")
	if err != nil {
		panic(err)
	}

	phone, err := values.NewPhoneNumber("+1234567890")
	if err != nil {
		panic(err)
	}

	address, err := values.NewAddress("123 Main St", "Springfield", "12345")
	if err != nil {
		panic(err)
	}

	dob, err := values.NewDateOfBirth(time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		panic(err)
	}

	// Create user entities
	user := entities.NewUser(name, email, phone, address, dob)

	// Print user details
	fmt.Println(user)
}
