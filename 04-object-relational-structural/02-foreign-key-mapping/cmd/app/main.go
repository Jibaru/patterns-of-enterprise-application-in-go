package main

import (
	"fmt"
	"log"

	"github.com/jibaru/foreign-key-mapping/internal/db"
	"github.com/jibaru/foreign-key-mapping/internal/domain"
)

func main() {
	// Initialize the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("failed to set up database: %v", err)
	}

	// Create repositories
	artistRepo := domain.NewArtistRepository(database)
	albumRepo := domain.NewAlbumRepository(database)

	// Create a new artist
	artist := &domain.Artist{Name: "The Beatles"}
	err = artistRepo.Insert(artist)
	if err != nil {
		log.Fatalf("failed to insert artist: %v", err)
	}
	fmt.Printf("Artist created: %+v\n", artist)

	// Create a few albums for the artist
	albums := []*domain.Album{
		{Title: "Abbey Road", Artist: artist},
		{Title: "Let It Be", Artist: artist},
	}

	for _, album := range albums {
		err = albumRepo.Insert(album)
		if err != nil {
			log.Fatalf("failed to insert album: %v", err)
		}
		fmt.Printf("Album created: %+v\n", album)
	}

	// Retrieve all albums for the artist
	retrievedAlbums, err := albumRepo.GetByArtist(artist.ID)
	if err != nil {
		log.Fatalf("failed to retrieve albums: %v", err)
	}

	fmt.Printf("Albums by %s:\n", artist.Name)
	for _, album := range retrievedAlbums {
		fmt.Printf("- %s\n", album.Title)
	}
}
