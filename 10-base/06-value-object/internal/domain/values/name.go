package values

import (
	"errors"
)

type Name struct {
	firstName string
	lastName  string
}

// NewName creates a new Name value object with validation
func NewName(firstName, lastName string) (*Name, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("invalid name details")
	}
	return &Name{firstName: firstName, lastName: lastName}, nil
}

func (n *Name) FirstName() string {
	return n.firstName
}

func (n *Name) LastName() string {
	return n.lastName
}
