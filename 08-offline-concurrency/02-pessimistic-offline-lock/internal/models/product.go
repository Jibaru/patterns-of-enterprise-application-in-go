package models

// Product represents a product entity in the database
type Product struct {
	ID     int
	Name   string
	Price  float64
	Locked bool
}
