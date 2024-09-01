// pricing_package.go
package pricing

type Pricing struct {
	TotalPrice float64
}

func (p *Pricing) SetPrice(price float64) {
	p.TotalPrice = price
}
