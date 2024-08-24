package repository

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/repository/internal/domain"
)

// UserRepository is the repository interface for User entity
type UserRepository interface {
	FindById(id int) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	Save(user *domain.User) error
	Update(user *domain.User) error
	Delete(id int) error
}

// userRepository is the concrete implementation of UserRepository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// FindById retrieves a user by their ID
func (r *userRepository) FindById(id int) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, first_name, last_name, email, age, is_active FROM users WHERE id = ?", id)

	var user domain.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	return &user, nil
}

// FindAll retrieves all users
func (r *userRepository) FindAll() ([]*domain.User, error) {
	rows, err := r.db.Query("SELECT id, first_name, last_name, email, age, is_active FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.IsActive)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// Save inserts a new user into the database
func (r *userRepository) Save(user *domain.User) error {
	result, err := r.db.Exec("INSERT INTO users (first_name, last_name, email, age, is_active) VALUES (?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Age, user.IsActive)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)

	return nil
}

// Update updates an existing user in the database
func (r *userRepository) Update(user *domain.User) error {
	_, err := r.db.Exec("UPDATE users SET first_name = ?, last_name = ?, email = ?, age = ?, is_active = ? WHERE id = ?",
		user.FirstName, user.LastName, user.Email, user.Age, user.IsActive, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a user from the database by ID
func (r *userRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
