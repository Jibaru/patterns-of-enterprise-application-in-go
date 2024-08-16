package main

import (
	"fmt"
	"log"

	"github.com/jibaru/dependent-mapping/internal/db"
	"github.com/jibaru/dependent-mapping/internal/domain"
)

func main() {
	// Initialize the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("failed to set up database: %v", err)
	}

	// Create an album with tracks
	album := &domain.Album{
		Title:  "The Dark Side of the Moon",
		Artist: "Pink Floyd",
		Tracks: []*domain.Track{
			{Title: "Speak to Me", Duration: 90},
			{Title: "Breathe", Duration: 163},
			{Title: "On the Run", Duration: 216},
		},
	}

	// Create AlbumMapper
	albumMapper := domain.NewAlbumMapper(database)

	// Insert the album along with its dependent tracks
	err = albumMapper.Insert(album)
	if err != nil {
		log.Fatalf("failed to insert album: %v", err)
	}
	fmt.Printf("Album inserted: %+v\n", album)

	// Retrieve the album and its tracks by ID
	retrievedAlbum, err := albumMapper.GetByID(album.ID)
	if err != nil {
		log.Fatalf("failed to retrieve album: %v", err)
	}

	fmt.Printf("Retrieved album: %s by %s\n", retrievedAlbum.Title, retrievedAlbum.Artist)
	for _, track := range retrievedAlbum.Tracks {
		fmt.Printf("- Track: %s, Duration: %d seconds\n", track.Title, track.Duration)
	}
}
