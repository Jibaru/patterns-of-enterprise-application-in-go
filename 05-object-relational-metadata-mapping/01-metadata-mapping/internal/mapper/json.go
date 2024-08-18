package mapper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// JSONMetaData holds the structure for mapping
type JSONMetaData struct {
	Table  string            `json:"table"`
	Fields map[string]string `json:"fields"`
}

// JSONMetadataMapper maps domain objects to the database based on metadata
type JSONMetadataMapper struct {
	DB       *sql.DB
	Metadata map[string]JSONMetaData
}

// NewJSONMetadataMapper loads metadata from a file and returns a MetadataMapper
func NewJSONMetadataMapper(db *sql.DB, metadataFile string) (*JSONMetadataMapper, error) {
	file, err := os.ReadFile(metadataFile)
	if err != nil {
		return nil, fmt.Errorf("could not read metadata file: %v", err)
	}

	var metadata map[string]JSONMetaData
	err = json.Unmarshal(file, &metadata)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal metadata: %v", err)
	}

	return &JSONMetadataMapper{DB: db, Metadata: metadata}, nil
}

// Save saves a domain object to the database using metadata
func (m *JSONMetadataMapper) Save(object interface{}) error {
	objectType := reflect.TypeOf(object).Elem().Name()
	objectValue := reflect.ValueOf(object).Elem()

	metaData, exists := m.Metadata[objectType]
	if !exists {
		return fmt.Errorf("metadata for object type %s not found", objectType)
	}

	fieldMap := metaData.Fields
	tableName := metaData.Table

	// Collect field values
	var columns []string
	var values []interface{}
	for domainField, dbColumn := range fieldMap {
		columns = append(columns, dbColumn)
		values = append(values, objectValue.FieldByName(domainField).Interface())
	}

	// Build the insert query
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, joinColumns(columns), joinPlaceholders(len(columns)))
	log.Printf("Executing from json mapper: %v; values: %v", query, values)

	// Execute query
	_, err := m.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("could not save %s: %v", objectType, err)
	}
	return nil
}

func joinColumns(columns []string) string {
	return "`" + strings.Join(columns, "`, `") + "`"
}

func joinPlaceholders(count int) string {
	placeholders := make([]string, count)
	for i := range placeholders {
		placeholders[i] = "?"
	}
	return strings.Join(placeholders, ", ")
}
