package label

import (
	"sort"

	"github.com/paulinethornberg/label-artist-fetcher/model"
)

type Builder struct {
}

func NewBuilder() *Builder {
	return &Builder{}
}

// TODO: WHY DO WE USE POINTERS HERE AGAIN?
func (b *Builder) LabelArtistCollection(songCollection model.SongCollection) []model.Label {
	// create map with map of label, artist collection.
	labelCollection := make(map[string][]model.Artist, 0)
	for _, song := range songCollection.Songs {
		if len(labelCollection[song.RecordLabel]) == 0 {
			labelCollection[song.RecordLabel] = []model.Artist{{Name: song.Artist}}
		} else {
			labelCollection[song.RecordLabel] = append(labelCollection[song.RecordLabel],
				model.Artist{
					Name: song.Artist},
			)
		}
	}

	// create slice from map
	labels := make([]model.Label, 0)
	for label, artists := range labelCollection {
		labels = append(labels, model.NewLabel(label, artists))
	}

	// sort alphabetically on label name
	sort.Slice(labels, func(i, j int) bool {
		return labels[i].LabelName < labels[j].LabelName
	})

	return labels
}
