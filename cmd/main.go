package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/paulinethornberg/label-artist-fetcher/builder"
	"github.com/paulinethornberg/label-artist-fetcher/builder/label"
	"github.com/paulinethornberg/label-artist-fetcher/config"
	"github.com/paulinethornberg/label-artist-fetcher/internalapi"
	"github.com/paulinethornberg/label-artist-fetcher/repository"
	"github.com/paulinethornberg/label-artist-fetcher/repository/sverigesradioopenapi"
)

var (
	labelRepository repository.Repository
	labelBuilder    builder.Builder
	internalAPI     *http.Server
)

func main() {
	setupEnvironment()
	setupRepository()
	setupBuilder()
	setupInternalAPI()

	AwaitShutdown(func() {
		log.Println("Received shutdown signal, stopping service")
		err := internalAPI.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	})
}

func setupEnvironment() {
	config.InternalAPIEndpoint = os.Getenv(config.InternalAPIEndpointKey)
	config.SverigesRadioEndpoint = os.Getenv(config.SverigesRadiondpointKey)
}

func setupBuilder() {
	labelBuilder = label.NewBuilder()
}

func setupRepository() {
	labelRepository = sverigesradioopenapi.NewRepository()
}

func setupInternalAPI() {
	router := mux.NewRouter()
	handler := internalapi.NewHandler(labelRepository, labelBuilder)
	router.HandleFunc("/labels", handler.GetLabels).Methods(http.MethodGet)
	go func() {
		log.Println("started internal API",
			"endpoint", config.InternalAPIEndpoint,
		)

		if err := http.ListenAndServe(config.InternalAPIEndpoint, router); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("error in internal API", "error", err)
		}

		log.Println("stopped internal API")
	}()
}

// AwaitShutdown blocks until a shutdown signal is received, at which point it will run the provided function
// if it is not nil.
func AwaitShutdown(fn func()) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	if fn != nil {
		fn()
	}

	os.Exit(0)
}
