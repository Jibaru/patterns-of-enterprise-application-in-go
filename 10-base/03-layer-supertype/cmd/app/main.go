package main

import (
	"encoding/json"
	"fmt"

	"github.com/jibaru/layer-supertype/internal/entities"
)

func main() {
	// Create a new Customer
	customer := entities.NewCustomer(1, "John Doe", "john.doe@example.com")
	data, _ := json.Marshal(customer)
	fmt.Printf("Customer: %v\n", string(data))

	// Create a new Order
	order := entities.NewOrder(1, "ORD-12345", 250.75)
	data, _ = json.Marshal(order)
	fmt.Printf("Order: %v\n", string(data))

	// Update customer timestamp
	customer.UpdateTimestamp()
	data, _ = json.Marshal(customer)
	fmt.Printf("Customer after timestamp update: %v\n", string(data))

	// Update Order timestamp
	order.UpdateTimestamp()
	data, _ = json.Marshal(order)
	fmt.Printf("Order after timestamp update: %v\n", string(data))
}
