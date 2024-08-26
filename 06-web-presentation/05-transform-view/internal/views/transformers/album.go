package transformers

import (
	"fmt"

	"github.com/jibaru/transform-view/internal/domain"
)

// TransformAlbum generates an HTML representation of an Album object.
func TransformAlbum(album domain.Album) string {
	return fmt.Sprintf(
		`<h1>%s</h1>
        <p>Artist: %s</p>`,
		album.Title,
		TransformArtist(album.Artist),
	)
}
