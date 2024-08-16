package main

import (
	"fmt"
	"log"

	"github.com/jibaru/identity-field/internal/db"
	"github.com/jibaru/identity-field/internal/domain"
)

func main() {
	// Initialize the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("failed to set up database: %v", err)
	}

	// Create a new device repository
	deviceRepo := domain.NewDeviceRepository(database)

	// Create a new smart device
	device := &domain.Device{
		Name:   "Smart Thermostat",
		Model:  "T1000",
		Status: "active",
	}

	// Insert the device into the database (ID will be automatically generated)
	err = deviceRepo.Insert(device)
	if err != nil {
		log.Fatalf("failed to insert device: %v", err)
	}

	// Display the device's identity field (ID)
	fmt.Printf("Device inserted with ID: %d\n", device.ID)

	// Retrieve the device by its ID
	retrievedDevice, err := deviceRepo.GetByID(device.ID)
	if err != nil {
		log.Fatalf("failed to retrieve device: %v", err)
	}

	fmt.Printf("Retrieved device: %+v\n", retrievedDevice)
}
