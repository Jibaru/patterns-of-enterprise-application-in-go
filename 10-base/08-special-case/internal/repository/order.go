package repository

import (
	"github.com/jibaru/special-case/internal/domain"
	"github.com/jibaru/special-case/internal/domain/specialcase"
)

// OrderRepository provides access to orders from storage
type OrderRepository struct {
	// In a real application, this might be a database connection
	orders map[int]domain.Order
}

// NewOrderRepository creates a new OrderRepository with some sample data
func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: map[int]domain.Order{
			1: {ID: 1, Items: []string{"item1", "item2"}},
			2: {ID: 2, Items: []string{"item3"}},
		},
	}
}

// FindOrderById retrieves an order by ID, returning a special case if not found
func (r *OrderRepository) FindOrderById(id int) domain.OrderInterface {
	if order, exists := r.orders[id]; exists {
		return order
	}
	return specialcase.NoOrder{}
}
