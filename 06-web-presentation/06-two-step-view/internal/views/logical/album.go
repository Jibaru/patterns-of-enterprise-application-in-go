package logical

import "github.com/jibaru/two-step-view/internal/domain"

type RenderedAlbum struct {
	Title  string
	Artist string
}

func RenderAlbum(album domain.Album) RenderedAlbum {
	return RenderedAlbum{
		Title:  album.Title,
		Artist: RenderArtist(album.Artist).Name,
	}
}
