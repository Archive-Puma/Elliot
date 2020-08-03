package server

import "html/template"

// NewTemplates TODO Doc
func NewTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}
