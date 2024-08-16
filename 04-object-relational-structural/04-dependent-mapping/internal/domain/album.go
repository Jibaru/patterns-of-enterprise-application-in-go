package domain

import (
	"database/sql"
	"fmt"
)

// Album represents a music album.
type Album struct {
	ID     int      // Identity field for the album
	Title  string   // Title of the album
	Artist string   // Artist of the album
	Tracks []*Track // Tracks associated with the album
}

// AlbumMapper is responsible for persisting albums and their dependent tracks.
type AlbumMapper struct {
	DB *sql.DB
}

// NewAlbumMapper creates a new AlbumMapper.
func NewAlbumMapper(db *sql.DB) *AlbumMapper {
	return &AlbumMapper{DB: db}
}

// Insert adds a new album to the database, including its dependent tracks.
func (mapper *AlbumMapper) Insert(album *Album) error {
	tx, err := mapper.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	query := `INSERT INTO albums (title, artist) VALUES (?, ?)`
	result, err := tx.Exec(query, album.Title, album.Artist)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert album: %w", err)
	}

	albumID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to retrieve album ID: %w", err)
	}
	album.ID = int(albumID)

	// Insert dependent tracks
	for _, track := range album.Tracks {
		err = mapper.insertTrack(tx, album.ID, track)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert track: %w", err)
		}
	}

	return tx.Commit()
}

// insertTrack adds a track to the database, linked to its album.
func (mapper *AlbumMapper) insertTrack(tx *sql.Tx, albumID int, track *Track) error {
	query := `INSERT INTO tracks (album_id, title, duration) VALUES (?, ?, ?)`
	_, err := tx.Exec(query, albumID, track.Title, track.Duration)
	if err != nil {
		return fmt.Errorf("failed to insert track: %w", err)
	}
	return nil
}

// GetByID retrieves an album and its associated tracks by the album's ID.
func (mapper *AlbumMapper) GetByID(id int) (*Album, error) {
	query := `SELECT id, title, artist FROM albums WHERE id = ?`
	row := mapper.DB.QueryRow(query, id)

	var album Album
	err := row.Scan(&album.ID, &album.Title, &album.Artist)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("album with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve album: %w", err)
	}

	// Load dependent tracks
	album.Tracks, err = mapper.getTracksForAlbum(album.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tracks: %w", err)
	}

	return &album, nil
}

// getTracksForAlbum retrieves all tracks associated with a specific album.
func (mapper *AlbumMapper) getTracksForAlbum(albumID int) ([]*Track, error) {
	query := `SELECT id, title, duration FROM tracks WHERE album_id = ?`

	rows, err := mapper.DB.Query(query, albumID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tracks: %w", err)
	}
	defer rows.Close()

	var tracks []*Track
	for rows.Next() {
		var track Track
		err := rows.Scan(&track.ID, &track.Title, &track.Duration)
		if err != nil {
			return nil, fmt.Errorf("failed to scan track: %w", err)
		}
		tracks = append(tracks, &track)
	}

	return tracks, nil
}
