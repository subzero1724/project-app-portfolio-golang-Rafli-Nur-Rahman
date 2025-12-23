package http

import (
	"html/template"
	"net/http"

	"project-app-portfolio-golang-rafli/internal/model"
)

type ProjectHandler struct {
	tmpl *template.Template
}

func NewProjectHandler(tmpl *template.Template) *ProjectHandler {
	return &ProjectHandler{tmpl: tmpl}
}

func (h *ProjectHandler) List(w http.ResponseWriter, r *http.Request) {
	data := model.PageData{
		Title: "Projects | Portfolio",
		Profile: model.Profile{
			Name: "Rafli Nur Rahman",
			Role: "Backend Developer | Golang | PostgreSQL",
		},
		Projects: []model.Project{
			{Title: "Portfolio Website", Description: "Website portfolio menggunakan Golang"},
			{Title: "Inventory CLI", Description: "Aplikasi inventaris berbasis CLI"},
		},
	}

	err := h.tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
