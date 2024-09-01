package processor

import (
	"errors"
	"testing"

	"github.com/jibaru/service-stub/internal/domain"
	"github.com/jibaru/service-stub/internal/services"
)

// TestUserProcessorWithStub tests UserProcessor using the UserServiceStub
func TestUserProcessorWithStub(t *testing.T) {
	// Create a stubbed UserService
	userService := services.NewUserServiceStub()

	// Create a UserProcessor with the stubbed service
	processor := NewUserProcessor(userService)

	// Define test cases
	tests := []struct {
		name          string
		userID        int
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name:          "Existing user",
			userID:        1,
			expectedUser:  &domain.User{ID: 1, Name: "Stub User", Age: 25},
			expectedError: nil,
		},
		{
			name:          "Non-existing user",
			userID:        2,
			expectedUser:  nil,
			expectedError: errors.New("user not found"),
		},
	}

	// Execute test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := processor.ProcessUserRequest(tt.userID)

			// Check if the returned user matches the expected user
			if user != nil && tt.expectedUser != nil {
				if user.ID != tt.expectedUser.ID || user.Name != tt.expectedUser.Name || user.Age != tt.expectedUser.Age {
					t.Errorf("Expected user %+v, got %+v", tt.expectedUser, user)
				}
			} else if user != tt.expectedUser {
				t.Errorf("Expected user %+v, got %+v", tt.expectedUser, user)
			}

			// Check if the returned error matches the expected error
			if err != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectedError, err)
				}
			}
		})
	}
}
