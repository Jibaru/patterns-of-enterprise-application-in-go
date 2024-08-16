package domain

import (
	"database/sql"
	"fmt"
)

// Device represents a smart device with an identity field (id).
type Device struct {
	ID     int    // Identity field
	Name   string // Name of the device
	Model  string // Model of the device
	Status string // Device status (e.g., "active", "inactive")
}

// DeviceRepository provides methods to interact with the devices in the database.
type DeviceRepository struct {
	DB *sql.DB
}

// NewDeviceRepository creates a new instance of DeviceRepository.
func NewDeviceRepository(db *sql.DB) *DeviceRepository {
	return &DeviceRepository{DB: db}
}

// Insert adds a new device to the database and automatically sets its identity field (ID).
func (repo *DeviceRepository) Insert(device *Device) error {
	query := `INSERT INTO devices (name, model, status) VALUES (?, ?, ?)`
	result, err := repo.DB.Exec(query, device.Name, device.Model, device.Status)
	if err != nil {
		return fmt.Errorf("failed to insert device: %w", err)
	}

	// Retrieve the automatically generated ID and assign it to the device
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve device ID: %w", err)
	}
	device.ID = int(id)

	return nil
}

// GetByID retrieves a device by its identity field (ID).
func (repo *DeviceRepository) GetByID(id int) (*Device, error) {
	query := `SELECT id, name, model, status FROM devices WHERE id = ?`
	row := repo.DB.QueryRow(query, id)

	var device Device
	err := row.Scan(&device.ID, &device.Name, &device.Model, &device.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("device with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve device: %w", err)
	}

	return &device, nil
}
