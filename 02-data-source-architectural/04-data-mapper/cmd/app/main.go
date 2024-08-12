package main

import (
	"fmt"
	"log"

	"github.com/jibaru/data-mapper/internal/db"
	"github.com/jibaru/data-mapper/internal/domain"
	"github.com/jibaru/data-mapper/internal/mappers"
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

	// Create a new PersonMapper instance
	personMapper := mappers.NewPersonMapper(dbConn)

	// Insert the person into the database
	err = personMapper.Insert(person)
	if err != nil {
		log.Fatalf("Failed to insert person: %v", err)
	}
	fmt.Println("Person inserted successfully with ID:", person.ID)

	// Update the person's details
	person.NumberOfDependents = 4
	err = personMapper.Update(person)
	if err != nil {
		log.Fatalf("Failed to update person: %v", err)
	}
	fmt.Println("Person updated successfully!")

	// Calculate exemptions
	exemptions := person.GetExemption()
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
	taxableEarnings := person.GetTaxableEarnings(earnings)
	fmt.Printf("Person's taxable earnings: $%.2f\n", taxableEarnings)
}
