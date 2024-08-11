package gateways

import (
	"database/sql"
)

// Employee represents an employee in the system
type Employee struct {
	ID         int
	FirstName  string
	LastName   string
	Department string
}

// EmployeeGateway provides methods to interact with the employees table
type EmployeeGateway struct {
	db *sql.DB
}

// NewEmployeeGateway creates a new EmployeeGateway
func NewEmployeeGateway(db *sql.DB) *EmployeeGateway {
	return &EmployeeGateway{db: db}
}

// Insert inserts a new employee into the employees table
func (g *EmployeeGateway) Insert(firstName, lastName, department string) (int, error) {
	result, err := g.db.Exec("INSERT INTO employees (first_name, last_name, department) VALUES (?, ?, ?)", firstName, lastName, department)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// FindByID retrieves an employee by ID from the employees table
func (g *EmployeeGateway) FindByID(id int) (*Employee, error) {
	var employee Employee
	err := g.db.QueryRow("SELECT id, first_name, last_name, department FROM employees WHERE id = ?", id).Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Department)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &employee, nil
}

// UpdateDepartment updates the department of an employee in the employees table
func (g *EmployeeGateway) UpdateDepartment(id int, department string) error {
	_, err := g.db.Exec("UPDATE employees SET department = ? WHERE id = ?", department, id)
	return err
}

// Delete deletes an employee from the employees table by ID
func (g *EmployeeGateway) Delete(id int) error {
	_, err := g.db.Exec("DELETE FROM employees WHERE id = ?", id)
	return err
}
