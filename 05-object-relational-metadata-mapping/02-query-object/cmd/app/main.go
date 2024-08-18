package main

import (
	"encoding/json"
	"log"

	"github.com/jibaru/query-object/internal/db"
	"github.com/jibaru/query-object/internal/persistence"
)

func main() {
	// Set up the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer database.Close()

	productRepository := persistence.NewProductRepository(database)

	// Create a query object to filter products
	queries := []persistence.Query{
		{
			{
				Operator: persistence.GreaterThanEqualsOp,
				Field:    "price",
				Value:    100,
			},
			{
				Operator: persistence.IncludesOp,
				Field:    "name",
				Value:    "to",
			},
		},
		{
			{
				Operator: persistence.LowerThanOp,
				Field:    "price",
				Value:    20,
			},
		},
		{
			{
				Operator: persistence.EqualsOp,
				Field:    "price",
				Value:    40,
			},
		},
	}

	for _, query := range queries {
		data, _ := json.Marshal(query)
		log.Printf("Query: %v", string(data))

		// Execute the query
		products, err := productRepository.Query(query)
		if err != nil {
			log.Fatalf("Failed to execute query: %v", err)
		}

		data, _ = json.Marshal(products)

		log.Printf("Products: %v", string(data))
	}
}
