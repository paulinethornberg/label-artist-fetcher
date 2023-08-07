package sverigesradioopenapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/paulinethornberg/label-artist-fetcher/config"

	"github.com/paulinethornberg/label-artist-fetcher/model"
)

const (
	defaultSize = "100"
	jsonFormat  = "json"
)

type Repository struct {
	client *http.Client
}

func NewRepository() *Repository {
	client := new(http.Client)
	return &Repository{
		client: client,
	}
}

func (r *Repository) GetPlaylistByChannel(channel model.ChannelID, fromTime, toTime time.Time) (*model.SongCollection, error) {
	startDate := fromTime.Format(time.RFC3339)
	endDate := toTime.Format(time.RFC3339)

	path := fmt.Sprintf("/playlists/getplaylistbychannelid?id=%s&startdatetime=%s&enddatetime=%s&size=%s&format=%s",
		channel,
		startDate,
		endDate,
		defaultSize,
		jsonFormat)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", config.SverigesRadioEndpoint, path), nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var songCollection model.SongCollection
	if err := json.Unmarshal(b, &songCollection); err != nil {
		return nil, err
	}

	return &songCollection, nil
}
