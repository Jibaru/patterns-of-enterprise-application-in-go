package domain

// Person represents a person in the system with business logic
type Person struct {
	ID                 int
	LastName           string
	FirstName          string
	NumberOfDependents int
}

// GetExemption calculates the number of tax exemptions based on the number of dependents
func (p *Person) GetExemption() int {
	// Assume each dependent counts as one exemption, plus one for the person themselves
	return p.NumberOfDependents + 1
}

// IsFlaggedForAudit determines if the person should be flagged for audit
func (p *Person) IsFlaggedForAudit() bool {
	// Simple rule: flag for audit if the person has more than 4 dependents
	return p.NumberOfDependents > 4
}

// GetTaxableEarnings calculates the taxable earnings based on some earnings value
func (p *Person) GetTaxableEarnings(earnings float64) float64 {
	exemptionAmount := float64(p.GetExemption()) * 2000 // Assume each exemption reduces taxable income by $2000
	taxableEarnings := earnings - exemptionAmount
	if taxableEarnings < 0 {
		taxableEarnings = 0
	}
	return taxableEarnings
}
