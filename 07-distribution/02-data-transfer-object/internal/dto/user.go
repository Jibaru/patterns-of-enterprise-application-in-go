package dto

// UserDTO is a Data Transfer Object for the User entity
type UserDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
