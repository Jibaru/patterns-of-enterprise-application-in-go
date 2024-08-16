package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jibaru/serialized-lob/internal/db"
	"github.com/jibaru/serialized-lob/internal/domain"
)

func main() {
	// Initialize the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("failed to set up database: %v", err)
	}

	// Create a hierarchical Department structure
	parentDept := &domain.Department{Name: "Parent Department"}
	childDept1 := &domain.Department{Name: "Child Department 1"}
	childDept2 := &domain.Department{Name: "Child Department 2"}
	parentDept.Children = []*domain.Department{childDept1, childDept2}

	// Create a Customer entity with serialized departments
	customer := &domain.Customer{
		Name:        "Acme Corporation",
		Departments: []*domain.Department{parentDept},
	}

	// Create CustomerMapper
	customerMapper := domain.NewCustomerMapper(database)

	// Insert the customer into the database
	err = customerMapper.Insert(customer)
	if err != nil {
		log.Fatalf("failed to insert customer: %v", err)
	}
	fmt.Println("Customer inserted successfully")

	// Retrieve the customer by ID
	retrievedCustomer, err := customerMapper.GetByID(customer.ID)
	if err != nil {
		log.Fatalf("failed to retrieve customer: %v", err)
	}

	data, _ := json.Marshal(retrievedCustomer)

	fmt.Printf("Retrieved Customer: %v", string(data))
}
