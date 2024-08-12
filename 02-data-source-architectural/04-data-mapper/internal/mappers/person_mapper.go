package mappers

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/data-mapper/internal/domain"
)

// PersonMapper handles database operations for Person entities
type PersonMapper struct {
	db *sql.DB
}

// NewPersonMapper creates a new instance of PersonMapper
func NewPersonMapper(db *sql.DB) *PersonMapper {
	return &PersonMapper{db: db}
}

// Insert inserts a Person instance into the persons table
func (pm *PersonMapper) Insert(person *domain.Person) error {
	result, err := pm.db.Exec("INSERT INTO persons (first_name, last_name, number_of_dependents) VALUES (?, ?, ?)", person.FirstName, person.LastName, person.NumberOfDependents)
	if err != nil {
		return fmt.Errorf("failed to insert person: %w", err)
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}
	person.ID = int(insertID)
	return nil
}

// Update updates a Person's data in the persons table
func (pm *PersonMapper) Update(person *domain.Person) error {
	_, err := pm.db.Exec("UPDATE persons SET first_name = ?, last_name = ?, number_of_dependents = ? WHERE id = ?", person.FirstName, person.LastName, person.NumberOfDependents, person.ID)
	if err != nil {
		return fmt.Errorf("failed to update person: %w", err)
	}
	return nil
}

// FindByID retrieves a Person instance from the database by its ID
func (pm *PersonMapper) FindByID(id int) (*domain.Person, error) {
	row := pm.db.QueryRow("SELECT id, first_name, last_name, number_of_dependents FROM persons WHERE id = ?", id)
	person := &domain.Person{}
	err := row.Scan(&person.ID, &person.FirstName, &person.LastName, &person.NumberOfDependents)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to find person by ID: %w", err)
	}
	return person, nil
}
