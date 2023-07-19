package internalapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/paulinethornberg/label-artist-fetcher/builder"
	"github.com/paulinethornberg/label-artist-fetcher/model"
	"github.com/paulinethornberg/label-artist-fetcher/repository"
)

// Handler represents an API handler for status
type Handler struct {
	repository repository.Repository
	builder    builder.Builder
}

func NewHandler(repository repository.Repository, builder builder.Builder) *Handler {
	handler := &Handler{
		repository: repository,
		builder:    builder,
	}
	return handler
}

//TODO CONSIDER IF I WANT THIS OR NOT
//func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	log.Println("hello debugger")
//	switch r.Method {
//	case http.MethodGet:
//		h.GetLabels(w, r)
//	default:
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//	}
//}

func (h *Handler) GetLabels(w http.ResponseWriter, r *http.Request) {

	fromString := r.URL.Query().Get("from")
	fromInt, err := strconv.ParseInt(fromString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid `from` timestamp"))
		return
	}
	fromTime := time.Unix(fromInt, 0)

	toString := r.URL.Query().Get("to")
	toInt, err := strconv.ParseInt(toString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid `to` timestamp"))
		return
	}
	toTime := time.Unix(toInt, 0)
	fmt.Println(fromTime.String())
	fmt.Println(toTime.String())

	// Handle GET requests here
	inputChannel := r.URL.Query().Get("channel")
	if !model.IsValidChannel(inputChannel) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("channel not supported"))
		return
	}

	id := model.ConvertFromStringToID(inputChannel)
	data, err := h.repository.GetPlaylistByChannel(id)
	if err != nil {
		fmt.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting playlist"))
		return
	}
	collection := h.builder.LabelArtistCollection(*data)

	jsonCollection, _ := json.Marshal(collection)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCollection)
}
