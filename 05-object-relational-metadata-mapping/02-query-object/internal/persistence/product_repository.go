package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jibaru/query-object/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func criteriaToSqlCondition(c Criteria) string {
	op := ""
	switch c.Operator {
	case GreaterThanOp:
		op = ">"
	case GreaterThanEqualsOp:
		op = ">="
	case LowerThanOp:
		op = "<"
	case LowerThanEqualsOp:
		op = "<="
	case EqualsOp:
		op = "="
	case NotEqualsOp:
		op = "<>"
	case IncludesOp:
		op = "LIKE"
		return fmt.Sprintf("%v %v '%v%v%v'", c.Field, op, "%", c.Value, "%")
	}

	return fmt.Sprintf("%v %v '%v'", c.Field, op, c.Value)
}

func (r *ProductRepository) Query(criterias []Criteria) ([]domain.Product, error) {
	query := "SELECT id, name, price FROM products"

	var conditions []string
	for _, criteria := range criterias {
		conditions = append(conditions, criteriaToSqlCondition(criteria))
	}

	query += " WHERE " + strings.Join(conditions, " AND ")
	log.Printf("Executing SQL: %v", query)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}
