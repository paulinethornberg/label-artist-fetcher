package label

import (
	"testing"

	"github.com/paulinethornberg/label-artist-fetcher/model"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_LabelArtistCollection(t *testing.T) {
	builder := NewBuilder()
	label1 := "A"
	label2 := "B"
	songCollection := model.SongCollection{Songs: []model.Song{
		{
			Title:       "Aqua Aura",
			Artist:      "Fricky",
			Composer:    "Fricky_man",
			RecordLabel: label1,
		},
		{
			Title:       "Liv",
			Artist:      "Regularfantasy",
			Composer:    "any",
			RecordLabel: label2,
		},
		{
			Title:       "Tiff",
			Artist:      "DTiffany, ABC",
			Composer:    "T",
			RecordLabel: label1,
		},
	}}
	collection := builder.LabelArtistCollection(songCollection)
	assert.Equal(t, label1, collection[0].LabelName) // should be sorted alphabetically by label
	assert.Equal(t, 3, len(collection[0].Artists))   // should be sorted alphabetically by label
	assert.Equal(t, label2, collection[1].LabelName)
}
