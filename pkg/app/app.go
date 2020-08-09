package app

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cosasdepuma/elliot/pkg/config"
	"github.com/cosasdepuma/elliot/pkg/database"
	"github.com/cosasdepuma/elliot/pkg/server"
	"github.com/cosasdepuma/elliot/pkg/templates"
)

// === GLOBAL VARIABLES ===

var Elliot = struct {
	DB        *database.Database
	Server    *http.Server
	Templates *template.Template
}{
	DB:        database.NewDatabase(),
	Server:    server.NewServer(),
	Templates: templates.NewTemplates(),
}

// Start is the method that starts the execution of the backend according to the configuration data provided in the config package
func Start() {
	Elliot.Server.Handler = configureRouter()
	// Initialize the server
	go func() {
		fmt.Printf("🐹  Running on\thttp://%s:%d/\n", config.Host, config.Port)
		if err := Elliot.Server.ListenAndServe(); err != nil {
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
	Elliot.Server.Shutdown(ctx)
	os.Exit(0)
}
