package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jibaru/lazy-load/internal/db"
	"github.com/jibaru/lazy-load/internal/services"
)

func main() {
	// Initialize the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	defer dbConn.Close()

	// Initialize the CustomerService
	customerService := services.NewCustomerService(dbConn)

	// Load a customer lazily with their orders
	customer, err := customerService.GetCustomerByID(1)
	if err != nil {
		log.Fatalf("Failed to get customer: %v", err)
	}
	data, _ := json.Marshal(customer)
	fmt.Printf("Customer: %+v\n", string(data))

	// Access customer's orders lazily (this will load the orders only now)
	orders, err := customer.Orders()
	if err != nil {
		log.Fatalf("Failed to get customer's orders: %v", err)
	}
	data, _ = json.Marshal(orders)
	fmt.Printf("Orders: %+v\n", string(data))

	// Access order lines for the first order (Lazy Load of order lines)
	if len(orders) > 0 {
		orderLines, err := orders[0].OrderLines()
		if err != nil {
			log.Fatalf("Failed to get order lines: %v", err)
		}
		data, _ = json.Marshal(orderLines)
		fmt.Printf("Order Lines for Order ID %d: %+v\n", orders[0].ID, string(data))
	}
}
