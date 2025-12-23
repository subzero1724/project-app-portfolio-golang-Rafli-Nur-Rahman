package http

import (
	"html/template"
	"net/http"

	"project-app-portfolio-golang-rafli/internal/model"
)

type ContactHandler struct {
	tmpl *template.Template
}

func NewContactHandler(tmpl *template.Template) *ContactHandler {
	return &ContactHandler{tmpl: tmpl}
}

func (h *ContactHandler) Form(w http.ResponseWriter, r *http.Request) {
	data := model.PageData{
		Title: "Contact | Portfolio",
		Profile: model.Profile{
			Name: "Rafli Nur Rahman",
			Role: "Backend Developer | Golang | PostgreSQL",
		},
	}

	err := h.tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
