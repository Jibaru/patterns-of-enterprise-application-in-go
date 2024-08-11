package modules

import (
	"database/sql"
	"fmt"
)

// Table Module - Product
type ProductModule struct {
	db *sql.DB
}

// NewProduct creates a new instance of ProductModule
func NewProduct(db *sql.DB) *ProductModule {
	return &ProductModule{db: db}
}

// GetProductType retrieves the type of a product by its ID
func (pm *ProductModule) GetProductType(productID int) (string, error) {
	var productType string
	err := pm.db.QueryRow("SELECT type FROM products WHERE id = ?", productID).Scan(&productType)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve product type: %w", err)
	}
	return productType, nil
}
