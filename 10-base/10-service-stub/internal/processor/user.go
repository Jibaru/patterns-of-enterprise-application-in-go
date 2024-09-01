package processor

import (
	"errors"

	"github.com/jibaru/service-stub/internal/domain"
)

// UserProcessor handles business logic related to users
type UserProcessor struct {
	userService domain.UserService
}

// NewUserProcessor creates a new UserProcessor with the provided UserService
func NewUserProcessor(service domain.UserService) *UserProcessor {
	return &UserProcessor{userService: service}
}

// ProcessUserRequest processes a user request by retrieving user information
func (p *UserProcessor) ProcessUserRequest(userID int) (*domain.User, error) {
	user, err := p.userService.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
