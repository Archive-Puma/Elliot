package server

// === IMPORTS ===

import (
	"net/http"

	"github.com/gorilla/mux"
)

// === PUBLIC METHODS ===

// NewRouter generates a new structure associated with the Router
func NewRouter() *mux.Router {
	// Serves static files and initialize the Router
	static := http.FileServer(http.Dir("public/"))
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))
	return router
}
