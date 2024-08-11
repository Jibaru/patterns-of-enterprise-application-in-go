package main

import (
	"fmt"
	"log"

	"github.com/jibaru/active-record/internal/db"
	"github.com/jibaru/active-record/internal/domain"
)

func main() {
	// Set up the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to set up the database: %v", err)
	}
	defer dbConn.Close()

	// Create a new Person instance
	person := &domain.Person{
		FirstName:          "John",
		LastName:           "Doe",
		NumberOfDependents: 3,
	}

	// Insert the person into the database
	err = person.Insert(dbConn)
	if err != nil {
		log.Fatalf("Failed to insert person: %v", err)
	}
	fmt.Println("Person inserted successfully with ID:", person.ID)

	// Update the person's details
	person.NumberOfDependents = 4
	err = person.Update(dbConn)
	if err != nil {
		log.Fatalf("Failed to update person: %v", err)
	}
	fmt.Println("Person updated successfully!")

	// Calculate exemptions
	exemptions := person.Exemption()
	fmt.Printf("Person has %d exemptions\n", exemptions)

	// Check if the person is flagged for audit
	flaggedForAudit := person.IsFlaggedForAudit()
	if flaggedForAudit {
		fmt.Println("Person is flagged for audit")
	} else {
		fmt.Println("Person is not flagged for audit")
	}

	// Calculate taxable earnings
	earnings := 50000.00 // Assume earnings
	taxableEarnings := person.TaxableEarnings(earnings)
	fmt.Printf("Person's taxable earnings: $%.2f\n", taxableEarnings)
}
