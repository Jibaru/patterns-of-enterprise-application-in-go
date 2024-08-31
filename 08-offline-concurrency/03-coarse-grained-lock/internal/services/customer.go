package services

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/jibaru/coarse-grained-lock/internal/models"
)

// CustomerService provides operations related to customers
type CustomerService struct {
	db   *sql.DB
	lock sync.Mutex // Coarse-grained lock
}

// NewCustomerService creates a new CustomerService
func NewCustomerService(db *sql.DB) *CustomerService {
	return &CustomerService{db: db}
}

// CreateCustomer creates a new customer in the database
func (s *CustomerService) CreateCustomer(name string) (*models.Customer, error) {
	result, err := s.db.Exec("INSERT INTO customers (name) VALUES (?)", name)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.Customer{
		ID:   int(id),
		Name: name,
	}, nil
}

// AddAddress adds a new address for the customer
func (s *CustomerService) AddAddress(customerID int, street, city, zip string) (*models.Address, error) {
	result, err := s.db.Exec("INSERT INTO addresses (customer_id, street, city, zip) VALUES (?, ?, ?, ?)", customerID, street, city, zip)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.Address{
		ID:         int(id),
		CustomerID: customerID,
		Street:     street,
		City:       city,
		Zip:        zip,
	}, nil
}

// UpdateAddressWithLock updates an address with a coarse-grained lock
func (s *CustomerService) UpdateAddressWithLock(customerID, addressID int, street, city, zip string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	address, err := s.GetAddress(addressID)
	if err != nil {
		return err
	}

	if address.CustomerID != customerID {
		return errors.New("address does not belong to the customer")
	}

	_, err = s.db.Exec("UPDATE addresses SET street = ?, city = ?, zip = ? WHERE id = ?", street, city, zip, addressID)
	if err != nil {
		return err
	}

	fmt.Printf("Updated Address: %+v\n", address)
	return nil
}

// GetAddress retrieves an address from the database
func (s *CustomerService) GetAddress(addressID int) (*models.Address, error) {
	row := s.db.QueryRow("SELECT id, customer_id, street, city, zip FROM addresses WHERE id = ?", addressID)

	var address models.Address
	if err := row.Scan(&address.ID, &address.CustomerID, &address.Street, &address.City, &address.Zip); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("address not found")
		}
		return nil, err
	}

	return &address, nil
}

// GetAddressOfCustomer retrieves addresses of customer from the database
func (s *CustomerService) GetAddressesOfCustomer(customerID int) ([]models.Address, error) {
	rows, err := s.db.Query("SELECT id, customer_id, street, city, zip FROM addresses WHERE customer_id = ?", customerID)
	if err != nil {
		return nil, err
	}

	var addresses []models.Address
	for rows.Next() {
		var address models.Address
		if err := rows.Scan(&address.ID, &address.CustomerID, &address.Street, &address.City, &address.Zip); err != nil {
			return nil, err
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}
