package services

// EmailService is a concrete implementation of the Service interface
type EmailService struct{}

// Execute runs the EmailService logic
func (e *EmailService) Execute() string {
	return "Email Service Executed"
}
