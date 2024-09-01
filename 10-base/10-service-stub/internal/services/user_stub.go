package services

import (
	"errors"

	"github.com/jibaru/service-stub/internal/domain"
)

// UserServiceStub is a stub implementation of the UserService interface for testing
type UserServiceStub struct{}

// NewUserServiceStub creates a new instance of UserServiceStub
func NewUserServiceStub() domain.UserService {
	return &UserServiceStub{}
}

// GetUserByID returns a hardcoded user or nil for testing purposes
func (s *UserServiceStub) GetUserByID(id int) (*domain.User, error) {
	// Return a predefined user for testing
	if id == 1 {
		return &domain.User{ID: 1, Name: "Stub User", Age: 25}, nil
	}
	return nil, errors.New("user not found")
}
