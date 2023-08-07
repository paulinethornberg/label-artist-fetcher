package sverigesradioopenapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/paulinethornberg/label-artist-fetcher/model"
)

const (
	baseURL     = "http://api.sr.se/api/v2"
	defaultSize = "100"
	jsonFormat  = "json"
	//startDate   = "2023-07-03T10:30:00Z" // todo have dates as input format
	//endDate     = "2023-07-04T17:30:00Z" // todo have dates as input format
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

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", baseURL, path), nil)
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
