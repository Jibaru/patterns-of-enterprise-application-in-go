package transformers

import (
	"fmt"

	"github.com/jibaru/transform-view/internal/domain"
)

// TransformArtist generates an HTML representation of an Artist object.
func TransformArtist(artist domain.Artist) string {
	return fmt.Sprintf(
		`<span>%s</span>`,
		artist.Name,
	)
}
