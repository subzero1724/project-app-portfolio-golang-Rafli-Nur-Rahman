package router

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"

	handler "project-app-portfolio-golang-rafli/internal/handler/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	homeHandler := handler.NewHomeHandler(tmpl)
	projectHandler := handler.NewProjectHandler(tmpl)
	contactHandler := handler.NewContactHandler(tmpl)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", homeHandler.Home)
	r.Get("/projects", projectHandler.List)
	r.Get("/contact", contactHandler.Form)

	return r
}
