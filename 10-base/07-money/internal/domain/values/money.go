package values

import (
	"errors"
	"fmt"
)

// Money represents a monetary value with a specific currency
type Money struct {
	amount   int64
	currency string
}

// NewMoney creates a new Money value object with validation
func NewMoney(amount int64, currency string) (*Money, error) {
	if amount < 0 {
		return nil, errors.New("amount cannot be negative")
	}
	if currency == "" {
		return nil, errors.New("currency cannot be empty")
	}
	return &Money{amount: amount, currency: currency}, nil
}

// Amount returns the amount of money
func (m *Money) Amount() int64 {
	return m.amount
}

// Currency returns the currency of the money
func (m *Money) Currency() string {
	return m.currency
}

// Add adds another Money value to this one, if they have the same currency
func (m *Money) Add(other *Money) (*Money, error) {
	if m.currency != other.currency {
		return nil, errors.New("currencies do not match")
	}
	return NewMoney(m.amount+other.amount, m.currency)
}

// Subtract subtracts another Money value from this one, if they have the same currency
func (m *Money) Subtract(other *Money) (*Money, error) {
	if m.currency != other.currency {
		return nil, errors.New("currencies do not match")
	}
	if m.amount < other.amount {
		return nil, errors.New("insufficient funds")
	}
	return NewMoney(m.amount-other.amount, m.currency)
}

// Multiply multiplies the Money value by a factor
func (m *Money) Multiply(factor int64) (*Money, error) {
	if factor < 0 {
		return nil, errors.New("factor cannot be negative")
	}
	return NewMoney(m.amount*factor, m.currency)
}

// Equals checks if two Money values are equal in amount and currency
func (m *Money) Equals(other *Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

// String returns the string representation of the Money value
func (m *Money) String() string {
	return fmt.Sprintf("%d %s", m.amount, m.currency)
}
