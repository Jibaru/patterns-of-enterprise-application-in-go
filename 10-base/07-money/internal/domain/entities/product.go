package entities

import (
	"fmt"

	"github.com/jibaru/money/internal/domain/values"
)

// Product represents an item with a price
type Product struct {
	name  string
	price *values.Money
}

// NewProduct creates a new Product with a name and price
func NewProduct(name string, price *values.Money) *Product {
	return &Product{name: name, price: price}
}

// Price returns the price of the product
func (p *Product) Price() *values.Money {
	return p.price
}

// String returns the string representation of the product
func (p *Product) String() string {
	return fmt.Sprintf("%s: %s", p.name, p.price)
}
