package mapper

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// TagMetadataMapper maps domain objects to the database using struct tags
type TagMetadataMapper struct {
	DB *sql.DB
}

// NewTagMetadataMapper returns a new instance of TagMapper
func NewTagMetadataMapper(db *sql.DB) *TagMetadataMapper {
	return &TagMetadataMapper{DB: db}
}

// Save saves a domain object to the database by reading struct tags
func (m *TagMetadataMapper) Save(object interface{}) error {
	objectType := reflect.TypeOf(object).Elem()
	objectValue := reflect.ValueOf(object).Elem()

	// Get the table name from the struct type (e.g., "Product" -> "products")
	tableName := strings.ToLower(objectType.Name()) + "s"

	var columns []string
	var values []interface{}
	var placeholders []string

	// Loop through the struct fields
	for i := 0; i < objectType.NumField(); i++ {
		field := objectType.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			continue // Skip fields without the 'db' tag
		}
		columns = append(columns, tag)
		values = append(values, objectValue.Field(i).Interface())
		placeholders = append(placeholders, "?")
	}

	// Build the insert query
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	log.Printf("Executing from tag mapper: %v; values: %v", query, values)

	// Execute query
	_, err := m.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("could not save %s: %v", objectType.Name(), err)
	}
	return nil
}
