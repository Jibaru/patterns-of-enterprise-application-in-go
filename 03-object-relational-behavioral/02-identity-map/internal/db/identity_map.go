package db

import (
	"sync"

	"github.com/jibaru/identity-map/internal/domain"
)

// IdentityMap stores the loaded objects in memory to avoid multiple database queries for the same object
type IdentityMap struct {
	mu     sync.Mutex
	drones map[int]*domain.Drone
}

// NewIdentityMap creates a new instance of IdentityMap
func NewIdentityMap() *IdentityMap {
	return &IdentityMap{
		drones: make(map[int]*domain.Drone),
	}
}

// Add stores a drone in the Identity Map
func (im *IdentityMap) Add(drone *domain.Drone) {
	im.mu.Lock()
	defer im.mu.Unlock()
	im.drones[drone.ID] = drone
}

// Get retrieves a drone from the Identity Map by its ID
func (im *IdentityMap) Get(id int) (*domain.Drone, bool) {
	im.mu.Lock()
	defer im.mu.Unlock()
	drone, found := im.drones[id]
	return drone, found
}
