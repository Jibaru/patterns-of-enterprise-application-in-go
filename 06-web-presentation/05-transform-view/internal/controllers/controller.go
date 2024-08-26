package controllers

import (
	"net/http"

	"github.com/jibaru/transform-view/internal/domain"
	"github.com/jibaru/transform-view/internal/views/transformers"
)

func AlbumView(w http.ResponseWriter, r *http.Request) {
	// Normally, you'd load this data from a database.
	artist := domain.Artist{ID: 1, Name: "The Beatles"}
	album := domain.Album{ID: 1, Title: "Abbey Road", Artist: artist}

	// Transform the album into HTML.
	htmlContent := transformers.TransformAlbum(album)

	// Write the HTML to the response.
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(transformers.Head + htmlContent + transformers.Foot))
}
