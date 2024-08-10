package models

// Order domain model
type Order struct {
	ID         int
	Customer   *Customer
	Items      []*OrderItem
	TotalPrice float64
}

// OrderItem domain model
type OrderItem struct {
	ID       int
	Product  *Product
	Quantity int
	Price    float64
}

// AddItem adds an item to the order
func (o *Order) AddItem(product *Product, quantity int) {
	item := &OrderItem{
		Product:  product,
		Quantity: quantity,
		Price:    product.Price * float64(quantity),
	}
	o.Items = append(o.Items, item)
	o.TotalPrice += item.Price
}

// ApplyDiscount applies a discount to the total price of the order
func (o *Order) ApplyDiscount(discount float64) {
	o.TotalPrice -= discount
}
