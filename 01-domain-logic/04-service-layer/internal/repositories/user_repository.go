package repositories

import (
	"database/sql"
	"errors"

	"github.com/jibaru/service-layer/internal/data"
)

// UserRepository provides methods to interact with the user data in the database
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Save saves a new user to the database
func (r *UserRepository) Save(user *data.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, password, first_name, last_name) VALUES (?, ?, ?, ?)", user.Username, user.Password, user.FirstName, user.LastName)
	return err
}

// FindByUsername retrieves a user by their username
func (r *UserRepository) FindByUsername(username string) (*data.User, error) {
	var user data.User
	err := r.db.QueryRow("SELECT id, username, password, first_name, last_name FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
