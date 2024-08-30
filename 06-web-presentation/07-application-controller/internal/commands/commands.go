package commands

import (
	"fmt"

	"github.com/jibaru/application-controller/internal/models"
)

type DomainCommand interface {
	Run() interface{}
}

type GetUserCommand struct {
	UserID int
}

func (c *GetUserCommand) Run() interface{} {
	// Simulate getting a user from a database
	user := &models.User{
		ID:    c.UserID,
		Name:  fmt.Sprintf("John Doe - %v", c.UserID),
		Email: fmt.Sprintf("john.doe-%v@example.com", c.UserID),
	}
	return user
}
