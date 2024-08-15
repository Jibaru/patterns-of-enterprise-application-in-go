package domain

import (
	"database/sql"
	"fmt"
)

// Customer represents a customer in the system
type Customer struct {
	ID           int
	Name         string
	orders       []*Order
	db           *sql.DB
	ordersLoaded bool
}

// NewCustomer creates a new Customer object
func NewCustomer(id int, name string, db *sql.DB) *Customer {
	return &Customer{
		ID:           id,
		Name:         name,
		db:           db,
		ordersLoaded: false,
	}
}

// Orders lazily loads the customer's orders if not already loaded
func (c *Customer) Orders() ([]*Order, error) {
	if !c.ordersLoaded {
		var orders []*Order

		query := "SELECT id, total FROM orders WHERE customer_id = ?"
		rows, err := c.db.Query(query, c.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to load orders: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			var order Order
			if err := rows.Scan(&order.ID, &order.Total); err != nil {
				return nil, fmt.Errorf("failed to scan order: %w", err)
			}
			order.db = c.db // Pass the database connection
			orders = append(orders, &order)
		}

		c.orders = orders
		c.ordersLoaded = true
	}
	return c.orders, nil
}
