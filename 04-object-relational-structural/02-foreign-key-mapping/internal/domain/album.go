package domain

import (
	"database/sql"
	"fmt"
)

// Album represents a music album.
type Album struct {
	ID     int     // Identity field for the album
	Title  string  // Title of the album
	Artist *Artist // Reference to the artist (foreign key)
}

// AlbumRepository provides methods to interact with albums in the database.
type AlbumRepository struct {
	DB *sql.DB
}

// NewAlbumRepository creates a new AlbumRepository.
func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{DB: db}
}

// Insert adds a new album to the database and associates it with an artist via a foreign key.
func (repo *AlbumRepository) Insert(album *Album) error {
	if album.Artist == nil || album.Artist.ID == 0 {
		return fmt.Errorf("album must be associated with a valid artist")
	}

	query := `INSERT INTO albums (title, artist_id) VALUES (?, ?)`
	result, err := repo.DB.Exec(query, album.Title, album.Artist.ID)
	if err != nil {
		return fmt.Errorf("failed to insert album: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve album ID: %w", err)
	}

	album.ID = int(id)
	return nil
}

// GetByArtist retrieves all albums associated with a given artist by artist ID.
func (repo *AlbumRepository) GetByArtist(artistID int) ([]*Album, error) {
	query := `SELECT id, title FROM albums WHERE artist_id = ?`
	rows, err := repo.DB.Query(query, artistID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve albums: %w", err)
	}
	defer rows.Close()

	var albums []*Album
	for rows.Next() {
		var album Album
		err := rows.Scan(&album.ID, &album.Title)
		if err != nil {
			return nil, fmt.Errorf("failed to scan album: %w", err)
		}

		album.Artist = &Artist{ID: artistID}
		albums = append(albums, &album)
	}

	return albums, nil
}
