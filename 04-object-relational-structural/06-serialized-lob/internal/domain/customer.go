package domain

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
)

// Customer represents a customer entity with serialized departments.
type Customer struct {
	ID          int
	Name        string
	Departments []*Department // Serialized as BLOB
}

// CustomerMapper handles persistence and retrieval of Customer entities.
type CustomerMapper struct {
	DB *sql.DB
}

// NewCustomerMapper creates a new CustomerMapper.
func NewCustomerMapper(db *sql.DB) *CustomerMapper {
	return &CustomerMapper{DB: db}
}

// Insert adds a new Customer to the database, serializing the departments.
func (mapper *CustomerMapper) Insert(customer *Customer) error {
	// Serialize the departments to BLOB
	departmentsBlob, err := serializeDepartments(customer.Departments)
	if err != nil {
		return fmt.Errorf("failed to serialize departments: %w", err)
	}

	query := `INSERT INTO customers (name, departments) VALUES (?, ?)`
	result, err := mapper.DB.Exec(query, customer.Name, departmentsBlob)
	if err != nil {
		return fmt.Errorf("failed to insert customer: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to insert customer: %w", err)
	}

	customer.ID = int(id)

	return nil
}

// GetByID retrieves a Customer by its ID, deserializing the departments.
func (mapper *CustomerMapper) GetByID(id int) (*Customer, error) {
	query := `SELECT id, name, departments FROM customers WHERE id = ?`
	row := mapper.DB.QueryRow(query, id)

	var customer Customer
	var departmentsBlob []byte

	err := row.Scan(&customer.ID, &customer.Name, &departmentsBlob)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("customer with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve customer: %w", err)
	}

	// Deserialize the departments from BLOB
	customer.Departments, err = deserializeDepartments(departmentsBlob)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize departments: %w", err)
	}

	return &customer, nil
}

// serializeDepartments serializes a slice of Department objects to a BLOB.
func serializeDepartments(departments []*Department) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(departments); err != nil {
		return nil, fmt.Errorf("failed to serialize departments: %w", err)
	}
	return buffer.Bytes(), nil
}

// deserializeDepartments deserializes a BLOB back into a slice of Department objects.
func deserializeDepartments(data []byte) ([]*Department, error) {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	var departments []*Department
	if err := decoder.Decode(&departments); err != nil {
		return nil, fmt.Errorf("failed to deserialize departments: %w", err)
	}
	return departments, nil
}
