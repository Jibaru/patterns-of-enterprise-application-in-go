package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/jibaru/coarse-grained-lock/internal/db"
	"github.com/jibaru/coarse-grained-lock/internal/services"
)

func main() {
	// Initialize the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Initialize the customer service
	customerService := services.NewCustomerService(dbConn)

	// Create a customer and addresses
	customer, err := customerService.CreateCustomer("John Doe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created Customer: %+v\n", customer)

	// Create addresses
	address1, err := customerService.AddAddress(customer.ID, "123 Main St", "City A", "12345")
	if err != nil {
		log.Fatal(err)
	}

	address2, err := customerService.AddAddress(customer.ID, "456 Oak St", "City B", "67890")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added Address: %+v\n", address1)
	fmt.Printf("Added Address: %+v\n", address2)

	addresses := []struct {
		street string
		city   string
		zip    string
	}{
		{street: "123 Maple St", city: "City A", zip: "12345"},
		{street: "456 Pine St", city: "City B", zip: "67890"},
	}

	// Simulate concurrent updates
	var wg sync.WaitGroup
	wg.Add(len(addresses))

	for i, address := range addresses {
		go func(addr struct {
			street string
			city   string
			zip    string
		}) {
			defer wg.Done()

			if err := customerService.UpdateAddressWithLock(
				customer.ID,
				address1.ID,
				addr.street,
				addr.city,
				addr.zip,
			); err != nil {
				fmt.Printf("Update %v failed: %v\n", i, err)
			} else {
				fmt.Printf("Update %v succeeded\n", i)
			}
		}(address)
	}
	// Wait for both updates to complete
	wg.Wait()

	updatedAddresses, err := customerService.GetAddressesOfCustomer(customer.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Addresses for customer %v on database: %v\n", customer.ID, updatedAddresses)
}
