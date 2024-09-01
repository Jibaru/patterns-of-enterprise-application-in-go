package entities

type Customer struct {
	Entity // Embeds the Entity (Layer Supertype) struct
	Name   string
	Email  string
}

// NewCustomer creates a new Customer instance
func NewCustomer(id int, name, email string) *Customer {
	return &Customer{
		Entity: *NewEntity(id), // Initializes the embedded Entity
		Name:   name,
		Email:  email,
	}
}
