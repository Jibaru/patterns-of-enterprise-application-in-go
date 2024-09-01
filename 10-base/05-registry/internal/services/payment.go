package services

// PaymentService is a concrete implementation of the Service interface
type PaymentService struct{}

// Execute runs the PaymentService logic
func (p *PaymentService) Execute() string {
	return "Payment Service Executed"
}
