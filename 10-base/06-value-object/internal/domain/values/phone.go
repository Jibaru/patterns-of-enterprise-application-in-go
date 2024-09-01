package values

import (
	"errors"
	"regexp"
)

type PhoneNumber struct {
	number string
}

// NewPhoneNumber creates a new PhoneNumber value object with validation
func NewPhoneNumber(number string) (*PhoneNumber, error) {
	if !isValidPhoneNumber(number) {
		return nil, errors.New("invalid phone number")
	}
	return &PhoneNumber{number: number}, nil
}

func (p *PhoneNumber) Number() string {
	return p.number
}

func isValidPhoneNumber(number string) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	return re.MatchString(number)
}
