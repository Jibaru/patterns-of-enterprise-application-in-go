package services

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/identity-map/internal/db"
	"github.com/jibaru/identity-map/internal/domain"
)

// DroneService handles the business logic related to drones
type DroneService struct {
	dbConn      *sql.DB
	identityMap *db.IdentityMap
}

// NewDroneService creates a new DroneService
func NewDroneService(dbConn *sql.DB) *DroneService {
	return &DroneService{
		dbConn:      dbConn,
		identityMap: db.NewIdentityMap(),
	}
}

// RegisterDrone adds a new drone to the database and Identity Map
func (s *DroneService) RegisterDrone(name, status string) error {
	query := "INSERT INTO drones (name, status) VALUES (?, ?)"
	result, err := s.dbConn.Exec(query, name, status)
	if err != nil {
		return fmt.Errorf("failed to insert drone: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	// Add to Identity Map
	drone := &domain.Drone{
		ID:     int(id),
		Name:   name,
		Status: status,
	}
	s.identityMap.Add(drone)

	return nil
}

// GetDroneByID fetches a drone by ID, first checking the Identity Map
func (s *DroneService) GetDroneByID(id int) (*domain.Drone, error) {
	// Check if the drone is already in the Identity Map
	if drone, found := s.identityMap.Get(id); found {
		return drone, nil
	}

	// If not in the Identity Map, fetch from database
	drone := &domain.Drone{}
	query := "SELECT id, name, status FROM drones WHERE id = ?"
	err := s.dbConn.QueryRow(query, id).Scan(&drone.ID, &drone.Name, &drone.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to get drone by ID: %w", err)
	}

	// Add to Identity Map
	s.identityMap.Add(drone)

	return drone, nil
}
