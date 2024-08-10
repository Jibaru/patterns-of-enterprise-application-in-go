package storage

import (
	"database/sql"

	"github.com/jibaru/domain-model/internal/models"
)

// Save saves the order to the database
func SaveOrder(order *models.Order, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	result, err := tx.Exec("INSERT INTO orders (customer_id, total) VALUES (?, ?)", order.Customer.ID, order.TotalPrice)
	if err != nil {
		return err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	order.ID = int(orderID)

	for _, item := range order.Items {
		_, err := tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)", order.ID, item.Product.ID, item.Quantity, item.Price)
		if err != nil {
			return err
		}
	}

	return nil
}
