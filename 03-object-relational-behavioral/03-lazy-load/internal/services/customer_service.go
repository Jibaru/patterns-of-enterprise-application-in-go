package services

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/lazy-load/internal/domain"
)

// CustomerService handles business logic related to customers
type CustomerService struct {
	db *sql.DB
}

// NewCustomerService creates a new CustomerService
func NewCustomerService(db *sql.DB) *CustomerService {
	return &CustomerService{db: db}
}

// GetCustomerByID fetches a customer by ID (without loading orders immediately)
func (s *CustomerService) GetCustomerByID(id int) (*domain.Customer, error) {
	customer := &domain.Customer{}
	query := "SELECT id, name FROM customers WHERE id = ?"
	err := s.db.QueryRow(query, id).Scan(&customer.ID, &customer.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch customer: %w", err)
	}

	// Pass the database connection for lazy loading orders
	return domain.NewCustomer(customer.ID, customer.Name, s.db), nil
}
