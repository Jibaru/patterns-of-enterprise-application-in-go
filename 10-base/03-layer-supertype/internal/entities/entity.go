package entities

import (
	"time"
)

// Entity represents a Layer Supertype with common fields and methods
type Entity struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewEntity initializes a new Entity with default values
func NewEntity(id int) *Entity {
	return &Entity{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateTimestamp updates the UpdatedAt field with the current time
func (e *Entity) UpdateTimestamp() {
	e.UpdatedAt = time.Now()
}
