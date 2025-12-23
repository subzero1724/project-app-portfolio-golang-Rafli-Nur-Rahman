package http

import (
	"encoding/json"
	"html/template"
	"net/http"
	"portfolio/internal/dto"
	"portfolio/internal/service"
	"portfolio/internal/util"

	"go.uber.org/zap"
)

type ContactHandler struct {
	contactService *service.ContactService
	logger         *zap.Logger
}

func NewContactHandler(contactService *service.ContactService, logger *zap.Logger) *ContactHandler {
	return &ContactHandler{
		contactService: contactService,
		logger:         logger,
	}
}

func (h *ContactHandler) ShowContactForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/contact.html")
	if err != nil {
		h.logger.Error("Failed to parse template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title": "Portfolio - Contact",
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		h.logger.Error("Failed to execute template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *ContactHandler) SubmitContact(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate
	if err := util.ValidateStruct(&req); err != nil {
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	contact, err := h.contactService.CreateContact(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to create contact", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}
