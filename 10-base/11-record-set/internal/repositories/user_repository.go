package repositories

import (
	"database/sql"
)

// UserRepository provides methods to interact with the user data in the database
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUsername retrieves a user by their username
func (r *UserRepository) GetAll() (*sql.Rows, error) {
	return r.db.Query("SELECT id, username, first_name, last_name FROM users")
}
