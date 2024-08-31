package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jibaru/optimistic-offline-lock/internal/db"
	"github.com/jibaru/optimistic-offline-lock/internal/services"
)

func main() {
	// Initialize the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Initialize the product service
	productService := services.NewProductService(dbConn)

	// Create a new product
	product, err := productService.CreateProduct("Laptop", 1000.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created Product: %+v\n", product)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	for i, newPrice := range []float64{950.0, 900.0} {
		go func(price float64, id int) {
			prod, err := productService.GetProduct(id)
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(1 * time.Second)

			prod.Price = price
			if err := productService.UpdateProduct(prod); err != nil {
				log.Printf("Failed to update product %v: %v\n", i, err)
			} else {
				fmt.Printf("Product %v updated: %+v\n", i, prod)
			}
			wg.Done()
		}(newPrice, product.ID)
	}
	wg.Wait()
}
