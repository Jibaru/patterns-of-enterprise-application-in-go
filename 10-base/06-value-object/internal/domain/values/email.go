package values

import (
	"errors"
	"regexp"
)

type Email struct {
	address string
}

// NewEmail creates a new Email value object with validation
func NewEmail(address string) (*Email, error) {
	if !isValidEmail(address) {
		return nil, errors.New("invalid email address")
	}
	return &Email{address: address}, nil
}

func (e *Email) Address() string {
	return e.address
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
