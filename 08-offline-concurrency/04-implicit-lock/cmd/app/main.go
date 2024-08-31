package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/jibaru/implicit-lock/internal/db"
	"github.com/jibaru/implicit-lock/internal/services"
)

func main() {
	// Initialize the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Initialize the order service
	orderService := services.NewOrderService(dbConn)

	// Create a new order
	order, err := orderService.CreateOrder("Product A", 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created Order: %+v\n", order)

	quantities := []int{10, 20, 30, 40}
	wg := &sync.WaitGroup{}
	wg.Add(len(quantities))

	// Simulate a concurrent update that will automatically be handled by implicit locking
	for i, qty := range quantities {
		go func(quantity int) {
			defer wg.Done()
			if err := orderService.UpdateOrderQuantity(order.ID, quantity); err != nil {
				fmt.Printf("Transaction %v: failed: %v\n", i, err)
			} else {
				fmt.Printf("Transaction %v: succeeded\n", i)
			}
		}(qty)
	}

	wg.Wait()
}
