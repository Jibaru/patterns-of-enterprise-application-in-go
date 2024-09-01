package services

import (
	"errors"

	"github.com/jibaru/service-stub/internal/domain"
)

// UserService is a concrete implementation of the UserService interface
type UserService struct{}

// NewUserService creates a new instance of UserServiceImpl
func NewUserService() *UserService {
	return &UserService{}
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(id int) (*domain.User, error) {
	// Simulate fetching user data from a database or external API
	if id == 1 {
		return &domain.User{ID: 1, Name: "John Doe", Age: 30}, nil
	}
	return nil, errors.New("user not found")
}
