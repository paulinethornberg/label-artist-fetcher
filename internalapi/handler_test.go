package internalapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/paulinethornberg/label-artist-fetcher/model"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"

	mocks2 "github.com/paulinethornberg/label-artist-fetcher/builder/mocks"

	"github.com/paulinethornberg/label-artist-fetcher/repository/mocks"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	repository          *mocks.Repository
	builder             *mocks2.Builder
	router              *mux.Router
	recorder            *httptest.ResponseRecorder
	testSongCollection  model.SongCollection
	testLabelCollection []model.Label
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (h *HandlerTestSuite) SetupTest() {
	h.router = mux.NewRouter()

	h.repository = &mocks.Repository{}
	h.builder = &mocks2.Builder{}
	h.recorder = httptest.NewRecorder()
	handler := NewHandler(h.repository, h.builder)
	h.router.HandleFunc("/labels", handler.GetLabels).Methods(http.MethodGet)

	// setup test data
	h.testSongCollection = model.SongCollection{Songs: []model.Song{{
		Title:       "songsong",
		Artist:      "songartist",
		Composer:    "composer",
		RecordLabel: "label",
	}}}

	h.testLabelCollection = []model.Label{
		{
			LabelName: "label",
			Artists: []model.Artist{
				{Name: "songartist"},
			},
		},
	}

}

func (h *HandlerTestSuite) TestGetLabel_OK() {
	fromTimestamp := time.Date(2023, 01, 01, 01, 01, 01, 00, time.UTC)
	toTimestamp := time.Date(2023, 02, 02, 02, 01, 01, 00, time.UTC)
	channel := model.ChannelP3

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/labels?from=%d&to=%d&channel=%s", fromTimestamp.Unix(), toTimestamp.Unix(), channel),
		nil,
	)

	h.repository.On("GetPlaylistByChannel",
		model.ConvertFromStringToID(channel),
		fromTimestamp,
		toTimestamp).
		Return(&h.testSongCollection, nil)

	h.builder.On("LabelArtistCollection",
		h.testSongCollection).
		Return(h.testLabelCollection, nil)

	h.router.ServeHTTP(h.recorder, req)
	require.Equal(h.T(), fasthttp.StatusOK, h.recorder.Code)

	var resp []model.Label
	err := json.Unmarshal(h.recorder.Body.Bytes(), &resp)
	require.NoError(h.T(), err)
	require.Equal(h.T(), resp[0].LabelName, h.testLabelCollection[0].LabelName)
}

func (h *HandlerTestSuite) TestGetLabel_MissingTimestamp_BadRequest() {
	fromTimestamp := time.Date(2023, 01, 01, 01, 01, 01, 00, time.UTC)
	channel := model.ChannelP3

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/labels?from=%d&channel=%s", fromTimestamp.Unix(), channel),
		nil,
	)

	h.router.ServeHTTP(h.recorder, req)
	require.Equal(h.T(), fasthttp.StatusBadRequest, h.recorder.Code)
}
