package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter TODO Doc
func NewRouter() *mux.Router {
	static := http.FileServer(http.Dir("public/"))
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))
	return router
}
