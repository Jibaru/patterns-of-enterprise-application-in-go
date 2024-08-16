package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jibaru/embedded-value/internal/db"
	"github.com/jibaru/embedded-value/internal/domain"
)

func main() {
	// Initialize the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("failed to set up database: %v", err)
	}

	employment := &domain.Employment{
		Person: domain.Person{
			ID:   1,
			Name: "John Doe",
		},
		Period: domain.DateRange{
			Start: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
		},
		Salary: domain.Money{
			Amount:   50000.00,
			Currency: "USD",
		},
	}

	// Create EmploymentMapper
	employmentMapper := domain.NewEmploymentMapper(database)

	// Insert the employment into the database
	err = employmentMapper.Insert(employment)
	if err != nil {
		log.Fatalf("failed to insert employment: %v", err)
	}
	fmt.Println("Employment inserted successfully")

	// Retrieve the employment by ID
	retrievedEmployment, err := employmentMapper.GetByID(employment.ID)
	if err != nil {
		log.Fatalf("failed to retrieve employment: %v", err)
	}

	data, _ := json.Marshal(retrievedEmployment)

	fmt.Printf("Retrieved Employment: %v\n", string(data))
}
