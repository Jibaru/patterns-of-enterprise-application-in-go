package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jibaru/implicit-lock/internal/models"
)

// OrderService provides operations related to orders
type OrderService struct {
	db *sql.DB
}

// NewOrderService creates a new OrderService
func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{db: db}
}

// CreateOrder creates a new order in the database
func (s *OrderService) CreateOrder(productName string, quantity int) (*models.Order, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO orders (product_name, quantity) VALUES (?, ?)", productName, quantity)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	order := &models.Order{
		ID:          int(id),
		ProductName: productName,
		Quantity:    quantity,
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return order, nil
}

// UpdateOrderQuantity updates the quantity of an order, relying on implicit locks within the transaction
func (s *OrderService) UpdateOrderQuantity(orderID int, quantity int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	order, err := s.GetOrder(tx, orderID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE orders SET quantity = ? WHERE id = ?", quantity, orderID)
	if err != nil {
		return err
	}

	fmt.Printf("Updated Order: %+v\n", order)

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetOrder retrieves an order from the database within a transaction
func (s *OrderService) GetOrder(tx *sql.Tx, orderID int) (*models.Order, error) {
	row := tx.QueryRow("SELECT id, product_name, quantity FROM orders WHERE id = ?", orderID)

	var order models.Order
	if err := row.Scan(&order.ID, &order.ProductName, &order.Quantity); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	return &order, nil
}
