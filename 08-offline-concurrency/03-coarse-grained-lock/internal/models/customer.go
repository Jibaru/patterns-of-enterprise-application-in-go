package models

// Customer represents a customer entity in the database
type Customer struct {
	ID        int
	Name      string
	Addresses []*Address
}
