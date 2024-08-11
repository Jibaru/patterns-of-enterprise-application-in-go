package modules

import (
	"database/sql"
	"fmt"
	"time"
)

// Table Module - RevenueRecognition
type RevenueRecognitionModule struct {
	db *sql.DB
}

// NewRevenueRecognition creates a new instance of RevenueRecognitionModule
func NewRevenueRecognition(db *sql.DB) *RevenueRecognitionModule {
	return &RevenueRecognitionModule{db: db}
}

// Insert records a revenue recognition entry
func (rm *RevenueRecognitionModule) Insert(contractID int, amount float64, date time.Time) error {
	_, err := rm.db.Exec("INSERT INTO revenue_recognitions (contract_id, amount, date) VALUES (?, ?, ?)", contractID, amount, date)
	if err != nil {
		return fmt.Errorf("failed to insert revenue recognition: %w", err)
	}
	return nil
}

// RecognizedRevenue retrieves the recognized revenue for a contract up to a given date
func (rm *RevenueRecognitionModule) RecognizedRevenue(contractID int, asOfDate time.Time) (float64, error) {
	var total float64
	err := rm.db.QueryRow("SELECT SUM(amount) FROM revenue_recognitions WHERE contract_id = ? AND date <= ?", contractID, asOfDate).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve recognized revenue: %w", err)
	}
	return total, nil
}
