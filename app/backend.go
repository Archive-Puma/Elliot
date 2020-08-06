package elliot

// === IMPORTS ===

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/server"
	"github.com/cosasdepuma/elliot/app/server/database"
)

// === STRUCTURES ===

// SBackend is a container structure with references to the database, router, server and HTML templates
type SBackend struct {
	DB        *database.Database
	Router    *mux.Router
	Server    *http.Server
	Templates *template.Template
}

// === VARIABLES ===

// Backend is the global instantiation of the container structure with references to the database, router, server and HTML templates
var Backend = &SBackend{
	DB:        database.NewDatabase(),
	Router:    server.NewRouter(),
	Server:    server.NewServer(),
	Templates: server.NewTemplates(),
}

// === STRUCTURES METHODS ===

// Start is the method that starts the execution of the backend according to the configuration data provided in the config package
func (b *SBackend) Start() {
	// Configure the routes
	b.ConfigureRoutes()
	// Initialize the server
	go func() {
		fmt.Printf("🐹 Running on: http://%s:%d/\n", config.Host, config.Port)
		if err := b.Server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	// Wait Ctrl+C signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	// Smooth shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	b.Server.Shutdown(ctx)
	os.Exit(0)
}
