package values

import (
	"errors"
	"time"
)

type DateOfBirth struct {
	dob time.Time
}

// NewDateOfBirth creates a new DateOfBirth value object with validation
func NewDateOfBirth(dob time.Time) (*DateOfBirth, error) {
	if dob.After(time.Now()) {
		return nil, errors.New("date of birth cannot be in the future")
	}
	return &DateOfBirth{dob: dob}, nil
}

func (d *DateOfBirth) DOB() time.Time {
	return d.dob
}
