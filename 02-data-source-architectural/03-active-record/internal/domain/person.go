package domain

import (
	"database/sql"
	"fmt"
)

// Person represents a person in the system
type Person struct {
	ID                 int
	LastName           string
	FirstName          string
	NumberOfDependents int
}

// Insert inserts the Person instance into the persons table
func (p *Person) Insert(db *sql.DB) error {
	result, err := db.Exec("INSERT INTO persons (first_name, last_name, number_of_dependents) VALUES (?, ?, ?)", p.FirstName, p.LastName, p.NumberOfDependents)
	if err != nil {
		return fmt.Errorf("failed to insert person: %w", err)
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}
	p.ID = int(insertID)
	return nil
}

// Update updates the Person's data in the persons table
func (p *Person) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE persons SET first_name = ?, last_name = ?, number_of_dependents = ? WHERE id = ?", p.FirstName, p.LastName, p.NumberOfDependents, p.ID)
	if err != nil {
		return fmt.Errorf("failed to update person: %w", err)
	}
	return nil
}

// Exemption calculates the number of tax exemptions based on the number of dependents
func (p *Person) Exemption() int {
	// Assume each dependent counts as one exemption, plus one for the person themselves
	return p.NumberOfDependents + 1
}

// IsFlaggedForAudit determines if the person should be flagged for audit
func (p *Person) IsFlaggedForAudit() bool {
	// Simple rule: flag for audit if the person has more than 4 dependents
	return p.NumberOfDependents > 4
}

// TaxableEarnings calculates the taxable earnings based on some earnings value
func (p *Person) TaxableEarnings(earnings float64) float64 {
	exemptionAmount := float64(p.Exemption()) * 2000 // Assume each exemption reduces taxable income by $2000
	taxableEarnings := earnings - exemptionAmount
	if taxableEarnings < 0 {
		taxableEarnings = 0
	}
	return taxableEarnings
}
