package repository

import (
	"github.com/paulinethornberg/label-artist-fetcher/model"
	"time"
)

type Repository interface {
	GetPlaylistByChannel(channel model.ChannelID, fromTime, toTime time.Time) (*model.SongCollection, error)
}
