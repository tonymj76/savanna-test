package main

import (
	"github.com/tonymj76/savannah/handlers"
	"github.com/tonymj76/savannah/routes"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// init gets called before the main function
func init() {
	if err := godotenv.Load(); err != nil {
		log.Error("No .env file found create")
		os.Exit(1)
	}
}

func main() {
	handler, err := handlers.NewRestService(handlers.WithServices(), handlers.WithDBSetup())
	if err != nil {
		log.Fatalf("error setting up new rest server. Err: %w", err)
	}
	route := routes.SetRouter(handler)

	if err := route.Run(":9090"); err != nil {
		log.Fatal("failed to start up server")
	}
}
