package main

import (
	"context"
	"fmt"
	"github.com/tonymj76/savannah/config"
	"github.com/tonymj76/savannah/handlers"
	"github.com/tonymj76/savannah/monitor"
	"github.com/tonymj76/savannah/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// running the monitor server as a goroutine to check and update DB for every hour
	go monitor.MRepo(handler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.GetEnv("PORT", "9090")),
		Handler: route,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
