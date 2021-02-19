package main

import (
	"net/http"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	usefulserver "github.com/mbfr/exampleserver/pkg"
)

func main() {
	log.Info("Starting server")
	err := run()
	if err != nil {
		log.WithError(err).Panic()
	}
}

func run() error {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	templateDir, ok := os.LookupEnv("TEMPLATE_DIR")
	if !ok {
		templateDir = "/templates"
	}

	server := usefulserver.GetServer(port, templateDir)
	defer server.Close()

	err := server.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return errors.Wrap(err, "unexpected error serving")
		}
	}

	return nil
}
