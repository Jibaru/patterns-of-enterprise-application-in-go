package modules

import (
	"database/sql"
	"fmt"
	"time"
)

// ContractModule module
type ContractModule struct {
	db *sql.DB
}

// NewContract creates a new instance of ContractModule
func NewContract(db *sql.DB) *ContractModule {
	return &ContractModule{db: db}
}

// RecognizeRevenue recognizes revenue for a contract
func (cm *ContractModule) RecognizeRevenue(contractID int) error {
	var productType string
	var totalRevenue float64
	var dateSigned time.Time

	// Retrieve contract details
	err := cm.db.QueryRow("SELECT p.type, c.total_revenue, c.date_signed FROM contracts c JOIN products p ON c.product_id = p.id WHERE c.id = ?", contractID).Scan(&productType, &totalRevenue, &dateSigned)
	if err != nil {
		return fmt.Errorf("failed to retrieve contract details: %w", err)
	}

	// Logic to recognize revenue based on product type
	rm := NewRevenueRecognition(cm.db)
	switch productType {
	case "Software":
		// Recognize revenue over three installments
		installment := totalRevenue / 3
		rm.Insert(contractID, installment, dateSigned)
		rm.Insert(contractID, installment, dateSigned.Add(30*24*time.Hour))
		rm.Insert(contractID, installment, dateSigned.Add(60*24*time.Hour))
	case "Service":
		// Recognize revenue immediately
		rm.Insert(contractID, totalRevenue, dateSigned)
	default:
		return fmt.Errorf("unknown product type: %s", productType)
	}

	return nil
}
