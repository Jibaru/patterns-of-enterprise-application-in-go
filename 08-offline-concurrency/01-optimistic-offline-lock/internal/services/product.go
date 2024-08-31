package services

import (
	"database/sql"
	"errors"

	"github.com/jibaru/optimistic-offline-lock/internal/models"
)

// ProductService provides operations related to products
type ProductService struct {
	db *sql.DB
}

// NewProductService creates a new ProductService
func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db: db}
}

// CreateProduct creates a new product in the database
func (s *ProductService) CreateProduct(name string, price float64) (*models.Product, error) {
	result, err := s.db.Exec("INSERT INTO products (name, price, version) VALUES (?, ?, ?)", name, price, 1)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.Product{
		ID:      int(id),
		Name:    name,
		Price:   price,
		Version: 1,
	}, nil
}

// GetProduct retrieves a product from the database
func (s *ProductService) GetProduct(id int) (*models.Product, error) {
	row := s.db.QueryRow("SELECT id, name, price, version FROM products WHERE id = ?", id)

	var product models.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Version); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

// UpdateProduct updates a product in the database using optimistic offline locking
func (s *ProductService) UpdateProduct(product *models.Product) error {
	result, err := s.db.Exec("UPDATE products SET name = ?, price = ?, version = version + 1 WHERE id = ? AND version = ?", product.Name, product.Price, product.ID, product.Version)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("conflict detected: product has been modified by another transaction")
	}

	product.Version++
	return nil
}
