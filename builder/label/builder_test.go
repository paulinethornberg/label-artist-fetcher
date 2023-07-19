package label

import (
	"fmt"
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
			Artist:      "DTiffany",
			Composer:    "T",
			RecordLabel: label1,
		},
	}}
	output := builder.LabelArtistCollection(songCollection)
	fmt.Print(output)
	assert.Equal(t, label1, output[0].LabelName)
	// TODO ADD github.com/stretchr/testifytestif
}
