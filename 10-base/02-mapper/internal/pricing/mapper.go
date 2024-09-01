// pricing_mapper.go
package pricing

import "github.com/jibaru/mapper/internal/models"

type PricingMapper struct{}

func (pm *PricingMapper) CalculatePricing(customer *models.Customer, lease *models.Lease, asset *models.Asset) *Pricing {
	basePrice := lease.DurationMonths * int(lease.MonthlyRate)
	assetValueFactor := asset.Value * 0.05
	customerDiscount := float64(100 - customer.Age)

	total := float64(basePrice) + assetValueFactor - customerDiscount

	pricing := &Pricing{}
	pricing.SetPrice(total)
	return pricing
}
