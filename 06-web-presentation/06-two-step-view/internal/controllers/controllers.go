package controllers

import (
	"net/http"

	"github.com/jibaru/two-step-view/internal/domain"
	"github.com/jibaru/two-step-view/internal/views/format"
	"github.com/jibaru/two-step-view/internal/views/logical"
)

func AlbumView(w http.ResponseWriter, r *http.Request) {
	// Normally, you'd load this data from a database.
	artist := domain.Artist{Name: "The Beatles"}
	album := domain.Album{Title: "Abbey Road", Artist: artist}

	// First stage: Render logical presentation.
	renderedAlbum := logical.RenderAlbum(album)

	// Second stage: Render HTML.
	htmlContent := format.RenderAlbumToHTML(renderedAlbum)

	// Write the HTML to the response.
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(format.Head + htmlContent + format.Foot))
}
