package sverigesradioopenapi

import (
	"testing"

	"github.com/paulinethornberg/label-artist-fetcher/model"
)

func TestRepository_GetPlaylistByChannel(t *testing.T) {
	repository := NewRepository()
	channelID := model.ConvertFromStringToID("p3")
	repository.GetPlaylistByChannel(channelID)
	// TODO: FIX TEST.
}
