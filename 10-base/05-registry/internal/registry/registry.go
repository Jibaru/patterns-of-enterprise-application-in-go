// registry.go
package registry

import (
	"errors"

	"github.com/jibaru/registry/internal/services"
)

// Registry holds registered services
type Registry struct {
	services map[string]services.Service
}

// NewRegistry creates a new Registry instance
func NewRegistry() *Registry {
	return &Registry{
		services: make(map[string]services.Service),
	}
}

// Register adds a service to the registry
func (r *Registry) Register(name string, svc services.Service) {
	r.services[name] = svc
}

// Get retrieves a service from the registry by name
func (r *Registry) Get(name string) (services.Service, error) {
	svc, exists := r.services[name]
	if !exists {
		return nil, errors.New("service not found")
	}
	return svc, nil
}
