package domain

// OrderInterface defines the behavior expected from any order type
type OrderInterface interface {
	GetID() int
}

// Order represents a typical order in the system
type Order struct {
	ID    int
	Items []string
}

// GetID returns the ID of the order
func (o Order) GetID() int {
	return o.ID
}
