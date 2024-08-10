package main

import (
	"fmt"
	"log"

	"github.com/jibaru/domain-model/internal/db"
	"github.com/jibaru/domain-model/internal/models"
	"github.com/jibaru/domain-model/internal/models/storage"
)

func main() {
	// Set up the database
	db, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer db.Close()

	// Retrieve customer and products from the database
	var customer models.Customer
	db.QueryRow("SELECT id, name FROM customers WHERE name = 'John Doe'").Scan(&customer.ID, &customer.Name)

	var laptop models.Product
	db.QueryRow("SELECT id, name, price FROM products WHERE name = 'Laptop'").Scan(&laptop.ID, &laptop.Name, &laptop.Price)

	var mouse models.Product
	db.QueryRow("SELECT id, name, price FROM products WHERE name = 'Mouse'").Scan(&mouse.ID, &mouse.Name, &mouse.Price)

	// Create a new order
	order := &models.Order{Customer: &customer}

	// Add items to the order
	order.AddItem(&laptop, 1)
	order.AddItem(&mouse, 2)

	// Apply a discount to the order
	order.ApplyDiscount(50.00)

	// Save the order to the database
	err = storage.SaveOrder(order, db)
	if err != nil {
		log.Fatalf("Failed to save order: %v", err)
	}

	fmt.Printf("Order ID %d created with total price: %.2f\n", order.ID, order.TotalPrice)
}
