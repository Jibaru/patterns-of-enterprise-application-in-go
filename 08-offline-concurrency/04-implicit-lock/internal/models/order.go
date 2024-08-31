package models

// Order represents an order entity in the database
type Order struct {
	ID          int
	ProductName string
	Quantity    int
}
