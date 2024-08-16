package domain

import (
	"database/sql"
	"fmt"
)

// Artist represents a music artist.
type Artist struct {
	ID   int    // Identity field for the artist
	Name string // Name of the artist
}

// ArtistRepository provides methods to interact with artists in the database.
type ArtistRepository struct {
	DB *sql.DB
}

// NewArtistRepository creates a new ArtistRepository.
func NewArtistRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{DB: db}
}

// Insert adds a new artist to the database.
func (repo *ArtistRepository) Insert(artist *Artist) error {
	query := `INSERT INTO artists (name) VALUES (?)`
	result, err := repo.DB.Exec(query, artist.Name)
	if err != nil {
		return fmt.Errorf("failed to insert artist: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve artist ID: %w", err)
	}

	artist.ID = int(id)
	return nil
}

// GetByID retrieves an artist by their ID.
func (repo *ArtistRepository) GetByID(id int) (*Artist, error) {
	query := `SELECT id, name FROM artists WHERE id = ?`
	row := repo.DB.QueryRow(query, id)

	var artist Artist
	err := row.Scan(&artist.ID, &artist.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("artist with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve artist: %w", err)
	}

	return &artist, nil
}
