package builder

import "github.com/paulinethornberg/label-artist-fetcher/model"

type Builder interface {
	LabelArtistCollection(songCollection model.SongCollection) []model.Label
}
