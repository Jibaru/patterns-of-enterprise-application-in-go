package main

import (
	"fmt"
	"log"

	"github.com/jibaru/registry/internal/registry"
	"github.com/jibaru/registry/internal/services"
)

func main() {
	// Initialize the registry
	reg := registry.NewRegistry()

	// Register services
	reg.Register("email", &services.EmailService{})
	reg.Register("payment", &services.PaymentService{})

	// Retrieve and use the email service
	emailSvc, err := reg.Get("email")
	if err != nil {
		log.Fatalf("Error retrieving service: %v", err)
	}
	fmt.Println(emailSvc.Execute())

	// Retrieve and use the payment service
	paymentSvc, err := reg.Get("payment")
	if err != nil {
		log.Fatalf("Error retrieving service: %v", err)
	}
	fmt.Println(paymentSvc.Execute())
}
