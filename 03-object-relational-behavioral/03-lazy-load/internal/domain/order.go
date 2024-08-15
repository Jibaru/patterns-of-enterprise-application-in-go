package domain

import (
	"database/sql"
	"fmt"
)

// Order represents an order made by a customer
type Order struct {
	ID          int
	CustomerID  int
	Total       float64
	orderLines  []*OrderLine
	db          *sql.DB
	linesLoaded bool
}

// OrderLine represents a line item in an order
type OrderLine struct {
	ID          int
	OrderID     int
	ProductName string
	Quantity    int
}

// OrderLines lazily loads the order lines if not already loaded
func (o *Order) OrderLines() ([]*OrderLine, error) {
	if !o.linesLoaded {
		var orderLines []*OrderLine

		query := "SELECT id, product_name, quantity FROM order_lines WHERE order_id = ?"
		rows, err := o.db.Query(query, o.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to load order lines: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			var orderLine OrderLine
			if err := rows.Scan(&orderLine.ID, &orderLine.ProductName, &orderLine.Quantity); err != nil {
				return nil, fmt.Errorf("failed to scan order line: %w", err)
			}
			orderLines = append(orderLines, &orderLine)
		}

		o.orderLines = orderLines
		o.linesLoaded = true
	}
	return o.orderLines, nil
}
