package main

import (
	"fmt"

	"github.com/jibaru/money/internal/domain/entities"
	"github.com/jibaru/money/internal/domain/values"
)

func main() {
	// Create Money objects
	price1, err := values.NewMoney(5000, "USD")
	if err != nil {
		panic(err)
	}

	price2, err := values.NewMoney(3000, "USD")
	if err != nil {
		panic(err)
	}

	// Create products
	product1 := entities.NewProduct("Laptop", price1)
	product2 := entities.NewProduct("Headphones", price2)

	// Display products
	fmt.Println(product1)
	fmt.Println(product2)

	// Perform operations on Money objects
	totalPrice, err := price1.Add(price2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total Price: %s\n", totalPrice)

	discountedPrice, err := price1.Subtract(price2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Discounted Price: %s\n", discountedPrice)

	multipliedPrice, err := price2.Multiply(2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Price after multiplication: %s\n", multipliedPrice)
}
