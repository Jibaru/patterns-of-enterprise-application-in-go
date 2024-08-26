package logical

import "github.com/jibaru/two-step-view/internal/domain"

type RenderedArtist struct {
	Name string
}

func RenderArtist(artist domain.Artist) RenderedArtist {
	return RenderedArtist{
		Name: artist.Name,
	}
}
