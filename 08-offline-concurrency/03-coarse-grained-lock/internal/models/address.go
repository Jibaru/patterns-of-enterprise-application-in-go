package models

// Address represents an address entity in the database
type Address struct {
	ID         int
	CustomerID int
	Street     string
	City       string
	Zip        string
}
