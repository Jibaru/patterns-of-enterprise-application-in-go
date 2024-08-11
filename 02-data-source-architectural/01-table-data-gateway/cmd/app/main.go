package main

import (
	"fmt"
	"log"

	"github.com/jibaru/table-data-gateway/internal/db"
	"github.com/jibaru/table-data-gateway/internal/gateways"
)

func main() {
	// Set up the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to set up the database: %v", err)
	}
	defer dbConn.Close()

	// Initialize the EmployeeGateway
	employeeGateway := gateways.NewEmployeeGateway(dbConn)

	// Insert a new employee
	id, err := employeeGateway.Insert("John", "Doe", "Engineering")
	if err != nil {
		log.Fatalf("Failed to insert employee: %v", err)
	}
	fmt.Println("Employee inserted successfully!")

	// Retrieve the employee by ID
	employee, err := employeeGateway.FindByID(id)
	if err != nil {
		log.Fatalf("Failed to retrieve employee: %v", err)
	}
	fmt.Printf("Employee details: %v\n", employee)

	// Update the employee's department
	err = employeeGateway.UpdateDepartment(id, "Marketing")
	if err != nil {
		log.Fatalf("Failed to update employee: %v", err)
	}
	fmt.Println("Employee updated successfully!")

	// Delete the employee
	err = employeeGateway.Delete(id)
	if err != nil {
		log.Fatalf("Failed to delete employee: %v", err)
	}
	fmt.Println("Employee deleted successfully!")
}
