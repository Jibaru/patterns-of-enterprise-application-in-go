package services

import (
	"database/sql"
	"errors"

	"github.com/jibaru/data-transfer-object/internal/dto"
	"github.com/jibaru/data-transfer-object/internal/models"
)

// UserService provides operations related to users
type UserService struct {
	db *sql.DB
}

// NewUserService creates a new UserService
func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// CreateUser creates a new user in the database and returns the created user as a DTO
func (s *UserService) CreateUser(name, email string) (*dto.UserDTO, error) {
	result, err := s.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID:    int(id),
		Name:  name,
		Email: email,
	}, nil
}

// GetUser retrieves a user from the database and returns it as a DTO
func (s *UserService) GetUser(id int) (*dto.UserDTO, error) {
	row := s.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
