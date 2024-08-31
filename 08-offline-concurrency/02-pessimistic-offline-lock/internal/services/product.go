package services

import (
	"database/sql"
	"errors"

	"github.com/jibaru/pessimistic-offline-lock/internal/models"
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
	result, err := s.db.Exec("INSERT INTO products (name, price, locked) VALUES (?, ?, ?)", name, price, false)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.Product{
		ID:    int(id),
		Name:  name,
		Price: price,
	}, nil
}

// GetProduct retrieves a product from the database
func (s *ProductService) GetProduct(id int) (*models.Product, error) {
	row := s.db.QueryRow("SELECT id, name, price, locked FROM products WHERE id = ?", id)

	var product models.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Locked); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

// LockProduct locks a product for updates
func (s *ProductService) LockProduct(id int) error {
	product, err := s.GetProduct(id)
	if err != nil {
		return err
	}

	if product.Locked {
		return errors.New("product is already locked by another transaction")
	}

	_, err = s.db.Exec("UPDATE products SET locked = ? WHERE id = ?", true, id)
	if err != nil {
		return err
	}

	return nil
}

// UnlockProduct unlocks a product after updates
func (s *ProductService) UnlockProduct(id int) error {
	_, err := s.db.Exec("UPDATE products SET locked = ? WHERE id = ?", false, id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateProduct updates a product in the database
func (s *ProductService) UpdateProduct(product *models.Product) error {
	if _, err := s.db.Exec("UPDATE products SET name = ?, price = ?, locked = ? WHERE id = ?", product.Name, product.Price, false, product.ID); err != nil {
		return err
	}

	return nil
}
