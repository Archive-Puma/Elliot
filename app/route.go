package elliot

// === IMPORTS ===

import (
	"net/http"
)

// === STRUCTURES METHODS ===

// ConfigureRoutes appends the necessary routes to the router for the backend to work properly
func (b *SBackend) ConfigureRoutes() {

	// --- Web ---

	// Index / Loader route
	b.Router.HandleFunc("/", wLoader).Methods("GET")
	// /domain category
	rdomain := b.Router.PathPrefix("/domain").Subrouter()
	rdomain.Use(mDomain)
	// /domain/osint
	rdomain.HandleFunc("/", wDomain).Methods("GET", "POST")

	b.Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		Backend.Templates.ExecuteTemplate(w, "test", "It works!")
	})

	// Link the Router with the Server
	b.Server.Handler = b.Router
}

// === PRIVATE METHODS ===

func wLoader(w http.ResponseWriter, r *http.Request) {
	Backend.Templates.ExecuteTemplate(w, "loader", nil)
}

func wDomain(w http.ResponseWriter, r *http.Request) {
	Backend.Templates.ExecuteTemplate(w, "domain", Backend.DB.StoredData().Domain)
}
