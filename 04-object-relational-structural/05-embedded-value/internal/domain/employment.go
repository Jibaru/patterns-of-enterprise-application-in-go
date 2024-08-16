package domain

import (
	"database/sql"
	"fmt"
)

// Employment represents an employment entity that embeds value objects.
type Employment struct {
	ID     int       // Identity field for the employment
	Person Person    // Embedded Person value object
	Period DateRange // Embedded DateRange value object for employment period
	Salary Money     // Embedded Money value object for salary
}

// EmploymentMapper is responsible for persisting Employment entities.
type EmploymentMapper struct {
	DB *sql.DB
}

// NewEmploymentMapper creates a new EmploymentMapper.
func NewEmploymentMapper(db *sql.DB) *EmploymentMapper {
	return &EmploymentMapper{DB: db}
}

// Insert adds a new Employment to the database, including its embedded values.
func (mapper *EmploymentMapper) Insert(employment *Employment) error {
	query := `INSERT INTO employments (personId, personName, start, end, salaryAmount, salaryCurrency) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := mapper.DB.Exec(query, employment.Person.ID, employment.Person.Name, employment.Period.Start, employment.Period.End, employment.Salary.Amount, employment.Salary.Currency)
	if err != nil {
		return fmt.Errorf("failed to insert employment: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get id of employment: %w", err)
	}
	employment.ID = int(id)

	return nil
}

// GetByID retrieves an Employment by its ID, along with its embedded values.
func (mapper *EmploymentMapper) GetByID(id int) (*Employment, error) {
	query := `SELECT id, personId, personName, start, end, salaryAmount, salaryCurrency FROM employments WHERE id = ?`
	row := mapper.DB.QueryRow(query, id)

	var employment Employment

	err := row.Scan(
		&employment.ID,
		&employment.Person.ID,
		&employment.Person.Name,
		&employment.Period.Start,
		&employment.Period.End,
		&employment.Salary.Amount,
		&employment.Salary.Currency,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("employment with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve employment: %w", err)
	}

	return &employment, nil
}
