package server

// === IMPORTS ===

import "html/template"

// === PUBLIC METHODS ===

// NewTemplates generates a new structure associated with the HTML templates served on the website
func NewTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}
