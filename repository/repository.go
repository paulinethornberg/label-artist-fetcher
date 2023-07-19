package repository

import (
	"github.com/paulinethornberg/label-artist-fetcher/model"
)

type Repository interface {
	GetPlaylistByChannel(channel model.ChannelID) (*model.SongCollection, error)
}
