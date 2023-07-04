package main

import (
	"fmt"
	"log"

	"github.com/paulinethornberg/label-artist-fetcher/builder"
	"github.com/paulinethornberg/label-artist-fetcher/builder/label"
	"github.com/paulinethornberg/label-artist-fetcher/model"
	"github.com/paulinethornberg/label-artist-fetcher/repository"
	"github.com/paulinethornberg/label-artist-fetcher/repository/sverigesradioopenapi"
)

var (
	labelRepository repository.Repository
	labelBuilder    builder.Builder
)

func main() {
	labelRepository = sverigesradioopenapi.NewRepository()
	labelBuilder = label.NewBuilder()

	// TODO GET "p3" as parameter from API and DO INPUT VALIDATION. have
	id := "p3"
	if !model.IsValidChannel(id) {
		log.Fatal("non valid channel input")
		// return bad request
	}

	data, err := labelRepository.GetPlaylistByChannel(model.ConvertFromStringToID("p3"))
	if err != nil {
		log.Fatal(data) // TODO DO NOT DO FATAL HERE
	}
	collection := labelBuilder.LabelArtistCollection(*data)

	fmt.Println(collection)
	// hämtar ut en kanals låtlista för ett visst tidsintervall (helt ok med ett hårdkodat datumintervall,
	//hårdkoda max antal låtar till 100,
	//och en hårdkodad radiokanal, förslagsvis P3).
}
