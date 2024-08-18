package main

import (
	"log"

	"github.com/jibaru/metadata-mapping/internal/db"
	"github.com/jibaru/metadata-mapping/internal/domain"
	"github.com/jibaru/metadata-mapping/internal/mapper"
)

func main() {
	// Set up the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer database.Close()

	// Initialize the metadata mapper
	jsonMapper, err := mapper.NewJSONMetadataMapper(database, "./internal/metadata/mapping.json")
	if err != nil {
		log.Fatalf("Failed to load metadata: %v", err)
	}
	tagMapper := mapper.NewTagMetadataMapper(database)

	// Create and save a product
	product := &domain.Product{
		ID:    1,
		Name:  "Laptop",
		Price: 999.99,
	}
	err = jsonMapper.Save(product)
	if err != nil {
		log.Fatalf("Failed to save product: %v", err)
	}

	// Create and save a customer
	customer := &domain.Customer{
		ID:       1,
		FullName: "Jane Doe",
		Email:    "jane.doe@example.com",
	}
	err = jsonMapper.Save(customer)
	if err != nil {
		log.Fatalf("Failed to save customer: %v", err)
	}

	log.Println("Data with json mapper saved successfully")

	// Create and save a product
	product = &domain.Product{
		ID:    2,
		Name:  "PC",
		Price: 123.99,
	}
	err = tagMapper.Save(product)
	if err != nil {
		log.Fatalf("Failed to save product: %v", err)
	}

	// Create and save a customer
	customer = &domain.Customer{
		ID:       2,
		FullName: "Jhon Doe",
		Email:    "jhon.doe@example.com",
	}
	err = tagMapper.Save(customer)
	if err != nil {
		log.Fatalf("Failed to save customer: %v", err)
	}

	log.Println("Data with tag mapper saved successfully")
}
