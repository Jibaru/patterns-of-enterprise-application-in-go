package entities

type Order struct {
	Entity      // Embeds the Entity (Layer Supertype) struct
	OrderNumber string
	Amount      float64
}

// NewOrder creates a new Order instance
func NewOrder(id int, orderNumber string, amount float64) *Order {
	return &Order{
		Entity:      *NewEntity(id), // Initializes the embedded Entity
		OrderNumber: orderNumber,
		Amount:      amount,
	}
}
