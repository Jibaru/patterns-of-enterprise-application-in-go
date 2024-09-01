package entities

import (
	"fmt"

	"github.com/jibaru/value-object/internal/domain/values"
)

type User struct {
	name        *values.Name
	email       *values.Email
	phone       *values.PhoneNumber
	address     *values.Address
	dateOfBirth *values.DateOfBirth
}

// NewUser creates a new User instance with validated value objects
func NewUser(name *values.Name, email *values.Email, phone *values.PhoneNumber, address *values.Address, dateOfBirth *values.DateOfBirth) *User {
	return &User{
		name:        name,
		email:       email,
		phone:       phone,
		address:     address,
		dateOfBirth: dateOfBirth,
	}
}

func (u *User) String() string {
	return fmt.Sprintf(
		"User(Name: %s %s, Email: %s, Phone: %s, Address: %s, %s, %s, DateOfBirth: %s)",
		u.name.FirstName(),
		u.name.LastName(),
		u.email.Address(),
		u.phone.Number(),
		u.address.Street(),
		u.address.City(),
		u.address.ZipCode(),
		u.dateOfBirth.DOB().Format("2006-01-02"),
	)
}
