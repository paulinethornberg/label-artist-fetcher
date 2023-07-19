package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/paulinethornberg/label-artist-fetcher/config"

	"github.com/paulinethornberg/label-artist-fetcher/builder/label"
	"github.com/paulinethornberg/label-artist-fetcher/repository/sverigesradioopenapi"

	"github.com/paulinethornberg/label-artist-fetcher/internalapi"

	"github.com/paulinethornberg/label-artist-fetcher/builder"
	"github.com/paulinethornberg/label-artist-fetcher/repository"
)

var (
	labelRepository repository.Repository
	labelBuilder    builder.Builder
	internalAPI     *http.Server
)

// TODO: Create docker compose file, docker file, make it startable via docker.
// add API handling -> should be able to get with a few parameters.
// consider adding dates
// update README
// ADD BASIC DOCS FOR API IN README
// ADD HOW TO START IN README

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
}

func setupBuilder() {
	labelBuilder = label.NewBuilder()
}

func setupRepository() {
	labelRepository = sverigesradioopenapi.NewRepository()
}

func setupInternalAPI() {
	handler := internalapi.NewHandler(labelRepository, labelBuilder)
	mux := http.NewServeMux()
	mux.HandleFunc("/labels", handler.GetLabels)
	go func() {
		log.Println("started internal API",
			"endpoint", config.InternalAPIEndpoint,
		)

		if err := http.ListenAndServe(config.InternalAPIEndpoint, mux); err != http.ErrServerClosed {
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
