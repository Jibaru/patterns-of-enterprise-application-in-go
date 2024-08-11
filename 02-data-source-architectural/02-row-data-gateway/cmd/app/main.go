package main

import (
	"fmt"
	"log"

	"github.com/jibaru/row-data-gateway/internal/db"
	"github.com/jibaru/row-data-gateway/internal/gateways"
)

func main() {
	// Set up the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to set up the database: %v", err)
	}
	defer dbConn.Close()

	// Create a new EmployeeGateway instance for a specific employee
	employee := gateways.NewEmployeeGateway(dbConn, 0, "John", "Doe", "Engineering")

	// Insert the employee into the database
	err = employee.Insert()
	if err != nil {
		log.Fatalf("Failed to insert employee: %v", err)
	}
	fmt.Println("Employee inserted successfully with ID:", employee.ID)

	// Update the employee's department
	employee.Department = "Marketing"
	err = employee.Update()
	if err != nil {
		log.Fatalf("Failed to update employee: %v", err)
	}
	fmt.Println("Employee updated successfully!")

	// Find the employee by ID
	foundEmployee, err := gateways.FindEmployeeByID(dbConn, employee.ID)
	if err != nil {
		log.Fatalf("Failed to find employee: %v", err)
	}
	fmt.Printf("Found Employee: %+v\n", foundEmployee)

	// Delete the employee
	err = foundEmployee.Delete()
	if err != nil {
		log.Fatalf("Failed to delete employee: %v", err)
	}
	fmt.Println("Employee deleted successfully!")
}
