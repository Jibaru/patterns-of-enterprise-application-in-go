package specialcase

import "github.com/jibaru/special-case/internal/domain"

// NoOrder is a special case that represents the absence of an order
type NoOrder struct{}

// NewNoOrder creates a new instance of NoOrder
func NewNoOrder() domain.OrderInterface {
	return NoOrder{}
}

// GetID returns a default ID for the special case, typically 0 or a negative value
func (NoOrder) GetID() int {
	return 0
}
