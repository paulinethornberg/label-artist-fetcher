package sverigesradioopenapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/paulinethornberg/label-artist-fetcher/model"
)

const (
	baseURL     = "http://api.sr.se/api/v2"
	defaultSize = "100"
	jsonFormat  = "json"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetPlaylistByChannel(channel model.ChannelID) (*model.SongCollection, error) {
	// https://api.sr.se/api/documentation/v2/metoder/musik.html
	// channelid (obligatorisk) - listar "låt just nu" endast för angiven kanal. Kanalerna fås med metoden lista kanaler
	//startDateTime (default: dagens datum utan tid)
	//endDateTime (default: startDateTime + en dag)
	//size Sidstorlek (default: 20)
	//format xml|json|jsonp (default: xml)
	//path := fmt.Sprintf("/playlists/getplaylistbychannelid?id=%s&startdatetime=%s&enddatetime=%s&size=%s&format=%s",
	//	channel,
	//	time.Now().Add(-(time.Hour * 24 * 30)),
	//	time.Now().String(),
	//	defaultSize,
	//	jsonFormat)

	// TODO figure out if dates are UNIX format or some other format
	path := fmt.Sprintf("/playlists/getplaylistbychannelid?id=%s&size=%s&format=%s",
		channel,
		defaultSize,
		jsonFormat)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", baseURL, path), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		return nil, err
	}
	var songCollection model.SongCollection
	if err := json.Unmarshal(b, &songCollection); err != nil { // Parse []byte to go struct pointer
		return nil, err
	}

	fmt.Println(songCollection)
	return &songCollection, nil
}
