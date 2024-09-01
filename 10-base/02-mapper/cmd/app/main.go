// main.go
package main

import (
	"fmt"

	"github.com/jibaru/mapper/internal/models"
	"github.com/jibaru/mapper/internal/pricing"
)

func main() {
	customer := &models.Customer{Name: "John Doe", Age: 30}
	lease := &models.Lease{DurationMonths: 12, MonthlyRate: 200.0}
	asset := &models.Asset{ID: "A123", Value: 15000.0}

	mapper := &pricing.PricingMapper{}
	pricing := mapper.CalculatePricing(customer, lease, asset)

	fmt.Printf("Total Pricing: $%.2f\n", pricing.TotalPrice)
}
