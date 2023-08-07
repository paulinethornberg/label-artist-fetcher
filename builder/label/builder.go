package label

import (
	"sort"
	"strings"

	"github.com/paulinethornberg/label-artist-fetcher/model"
)

type Builder struct {
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) LabelArtistCollection(songCollection model.SongCollection) []model.Label {
	// create map with map of label, artist collection.
	labelCollection := make(map[string][]model.Artist)
	for _, song := range songCollection.Songs {
		if len(labelCollection[song.RecordLabel]) == 0 {
			// label is not yet added to map, thus add new slice with artists
			labelCollection[song.RecordLabel] = buildArrayOfArist(song.Artist)
			continue
		}
		// add artist to existing slice for label
		// TODO: improvement: this can now create duplicates - make sure to check if artists are already added to label
		labelCollection[song.RecordLabel] = append(labelCollection[song.RecordLabel], buildArrayOfArist(song.Artist)...)
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

// buildArrayOfArist handles the cases where a song is created by multiple artists
func buildArrayOfArist(artist string) []model.Artist {
	artistNames := strings.Split(artist, ",")
	artists := make([]model.Artist, 0)
	for _, a := range artistNames {
		artist := model.Artist{Name: a}
		artists = append(artists, artist)
	}
	return artists
}
