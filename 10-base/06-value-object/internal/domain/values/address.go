package values

import (
	"errors"
)

type Address struct {
	street  string
	city    string
	zipCode string
}

// NewAddress creates a new Address value object with validation
func NewAddress(street, city, zipCode string) (*Address, error) {
	if street == "" || city == "" || zipCode == "" {
		return nil, errors.New("invalid address details")
	}
	return &Address{street: street, city: city, zipCode: zipCode}, nil
}

func (a *Address) Street() string {
	return a.street
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) ZipCode() string {
	return a.zipCode
}
