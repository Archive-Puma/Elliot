package elliot

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
	"github.com/cosasdepuma/elliot/app/modules"
	"github.com/cosasdepuma/elliot/app/server"
)

type SBackend struct {
	DB        *server.DB
	Router    *mux.Router
	Server    *http.Server
	Templates *template.Template
}

// Backend TODO Doc
var Backend = &SBackend{
	DB:        server.NewDatabase(),
	Router:    server.NewRouter(),
	Server:    server.NewServer(),
	Templates: server.NewTemplates(),
}

// Entrypoint TODO Doc
func (b *SBackend) Start() {
	b.Router.HandleFunc("/", b.wIndex).Methods("GET")
	b.Router.HandleFunc("/dashboard", b.wDashboard).Methods("GET", "POST")
	b.Router.HandleFunc("/api/target", b.apiTarget).Methods("POST")

	b.Server.Handler = b.Router
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

// =================== WEB CONTENT ===================

func (b SBackend) wIndex(w http.ResponseWriter, r *http.Request) {
	b.Templates.ExecuteTemplate(w, "loader", nil)
}

func (b SBackend) wDashboard(w http.ResponseWriter, r *http.Request) {
	data := b.DB.GetAll()
	b.Templates.ExecuteTemplate(w, "dashboard", data)
}

// =================== API ===================

func (b SBackend) apiTarget(w http.ResponseWriter, r *http.Request) {
	target := r.FormValue("target")

	if len(target) > 0 {
		log.Println(fmt.Sprintf("Target added: %s", target))
		b.DB.SetTarget(target)

		cSubdomains := make(chan []string, 1)

		go modules.Subdomains(target, &cSubdomains)

		subdomains := <-cSubdomains
		if subdomains != nil {
			b.DB.SetSubdomains(subdomains)
		}
	}

	http.Redirect(w, r, "/dashboard", 301)
}
