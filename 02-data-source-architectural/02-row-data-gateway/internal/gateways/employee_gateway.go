package gateways

import (
	"database/sql"
)

// EmployeeGateway represents a single row in the employees table
type EmployeeGateway struct {
	db         *sql.DB
	ID         int
	FirstName  string
	LastName   string
	Department string
}

// NewEmployeeGateway creates a new EmployeeGateway instance
func NewEmployeeGateway(db *sql.DB, id int, firstName, lastName, department string) *EmployeeGateway {
	return &EmployeeGateway{
		db:         db,
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		Department: department,
	}
}

// Insert inserts the EmployeeGateway instance as a new row in the employees table
func (e *EmployeeGateway) Insert() error {
	result, err := e.db.Exec("INSERT INTO employees (first_name, last_name, department) VALUES (?, ?, ?)", e.FirstName, e.LastName, e.Department)
	if err != nil {
		return err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = int(insertID)
	return nil
}

// Update updates the current row in the employees table with the EmployeeGateway instance data
func (e *EmployeeGateway) Update() error {
	_, err := e.db.Exec("UPDATE employees SET first_name = ?, last_name = ?, department = ? WHERE id = ?", e.FirstName, e.LastName, e.Department, e.ID)
	return err
}

// Delete deletes the current row from the employees table
func (e *EmployeeGateway) Delete() error {
	_, err := e.db.Exec("DELETE FROM employees WHERE id = ?", e.ID)
	return err
}

// FindEmployeeByID retrieves an employee by ID and returns a populated EmployeeGateway instance
func FindEmployeeByID(db *sql.DB, id int) (*EmployeeGateway, error) {
	var employee EmployeeGateway
	err := db.QueryRow("SELECT id, first_name, last_name, department FROM employees WHERE id = ?", id).Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Department)
	if err != nil {
		return nil, err
	}
	employee.db = db
	return &employee, nil
}
