package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/jibaru/service-layer/internal/data"
	"github.com/jibaru/service-layer/internal/repositories"
)

// UserService provides business logic related to users
type UserService struct {
	userRepository *repositories.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

// RegisterUser registers a new user in the system
func (s *UserService) RegisterUser(username, password, firstName, lastName string) error {
	// Check if the username already exists
	existingUser, _ := s.userRepository.FindByUsername(username)
	if existingUser != nil {
		return errors.New("username already taken")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create the user
	user := &data.User{
		Username:  username,
		Password:  string(hashedPassword),
		FirstName: firstName,
		LastName:  lastName,
	}

	// Save the user in the repository
	return s.userRepository.Save(user)
}

// AuthenticateUser verifies the username and password
func (s *UserService) AuthenticateUser(username, password string) (bool, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

// GetUserDetails retrieves user details by username
func (s *UserService) GetUserDetails(username string) (*data.User, error) {
	return s.userRepository.FindByUsername(username)
}
