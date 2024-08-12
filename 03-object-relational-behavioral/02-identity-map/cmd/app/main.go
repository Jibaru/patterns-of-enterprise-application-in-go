package main

import (
	"fmt"
	"log"

	"github.com/jibaru/identity-map/internal/db"
	"github.com/jibaru/identity-map/internal/services"
)

func main() {
	// Initialize the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	defer dbConn.Close()

	// Initialize the DroneService
	droneService := services.NewDroneService(dbConn)

	// Register a new drone
	err = droneService.RegisterDrone("Drone A", "Available")
	if err != nil {
		log.Fatalf("Failed to register drone: %v", err)
	}

	// Fetch the same drone by ID
	drone, err := droneService.GetDroneByID(1)
	if err != nil {
		log.Fatalf("Failed to get drone by ID: %v", err)
	}
	fmt.Printf("Fetched drone: %+v\n", drone)

	// Fetch the drone again (should come from Identity Map)
	drone, err = droneService.GetDroneByID(1)
	if err != nil {
		log.Fatalf("Failed to get drone by ID: %v", err)
	}
	fmt.Printf("Fetched drone from Identity Map: %+v\n", drone)
}
