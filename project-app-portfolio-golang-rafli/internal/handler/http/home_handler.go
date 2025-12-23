package http

import (
	"html/template"
	"net/http"

	"project-app-portfolio-golang-rafli/internal/model"
)

type HomeHandler struct {
	tmpl *template.Template
}

func NewHomeHandler(tmpl *template.Template) *HomeHandler {
	return &HomeHandler{tmpl: tmpl}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	data := model.PageData{
		Title: "Home | Portfolio",
		Profile: model.Profile{
			Name:  "Rafli Nur Rahman",
			Role:  "Backend Developer | Golang | PostgreSQL",
			About: "Saya adalah backend developer yang berfokus pada Golang dan REST API.",
		},
	}

	err := h.tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
