package app

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/cosasdepuma/elliot/pkg/modules"
)

// === PRIVATE METHODS ===

func configureRouter() *mux.Router {
	// Serves static files and initialize the Router
	static := http.FileServer(http.Dir("public/"))
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))
	// Index / Loader route
	router.HandleFunc("/", wLoader).Methods("GET")
	// /domain category
	rdomain := router.PathPrefix("/domain").Subrouter()
	rdomain.Use(mDomain)
	rdomain.HandleFunc("", wDomain).Methods("GET", "POST")

	return router
}

// === Routes ===

func wLoader(w http.ResponseWriter, r *http.Request) {
	Elliot.Templates.ExecuteTemplate(w, "loader", nil)
}

func wDomain(w http.ResponseWriter, r *http.Request) {
	Elliot.Templates.ExecuteTemplate(w, "domain", Elliot.DB.StoredDomainData())
}

// === Middlewares ===

func mDomain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if domain := r.FormValue("domain"); len(domain) > 0 {
			Elliot.DB.Purge()
			modules.RunDomain(strings.TrimSpace(strings.ToLower(domain)), Elliot.DB)
		}
		next.ServeHTTP(w, r)
	})
}
