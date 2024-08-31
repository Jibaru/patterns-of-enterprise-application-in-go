package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jibaru/pessimistic-offline-lock/internal/db"
	"github.com/jibaru/pessimistic-offline-lock/internal/models"
	"github.com/jibaru/pessimistic-offline-lock/internal/services"
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
	createdProduct, err := productService.CreateProduct("Smartphone", 500.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created Product: %+v\n", createdProduct)

	newPrices := []float64{450.0, 400.0, 300.0, 320.0}

	wg := &sync.WaitGroup{}
	wg.Add(len(newPrices))

	// Simulate many transactions trying to update the product
	for i, newPrice := range newPrices {
		go func(price float64, product models.Product) {
			defer wg.Done()

			err := productService.LockProduct(product.ID)
			for err != nil {
				fmt.Printf("Transaction %v: Failed to lock product: %v\n", i, err)
				fmt.Printf("Transaction %v: Waiting to get product lock\n", i)
				time.Sleep(1 * time.Second)
				err = productService.LockProduct(product.ID)
			}
			fmt.Printf("Transaction %v: Got product lock\n", i)
			defer productService.UnlockProduct(product.ID)

			time.Sleep(1 * time.Second)

			product.Price = price
			if err := productService.UpdateProduct(&product); err != nil {
				fmt.Printf("Transaction %v: Failed to update product: %v\n", i, err)
			} else {
				fmt.Printf("Transaction %v: Updated Product: %+v\n", i, product)
			}
		}(newPrice, *createdProduct)
	}

	wg.Wait()
}
