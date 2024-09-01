package domain

// User represents a user in the system
type User struct {
	ID   int
	Name string
	Age  int
}

// UserService defines the interface for user-related operations
type UserService interface {
	// GetUserByID retrieves a user by their ID
	GetUserByID(id int) (*User, error)
}
