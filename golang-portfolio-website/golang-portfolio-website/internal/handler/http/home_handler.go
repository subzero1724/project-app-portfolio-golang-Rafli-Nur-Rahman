package http

import (
	"html/template"
	"net/http"
	"portfolio/internal/service"

	"go.uber.org/zap"
)

type HomeHandler struct {
	projectService *service.ProjectService
	userService    *service.UserService
	logger         *zap.Logger
}

func NewHomeHandler(projectService *service.ProjectService, userService *service.UserService, logger *zap.Logger) *HomeHandler {
	return &HomeHandler{
		projectService: projectService,
		userService:    userService,
		logger:         logger,
	}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	// Get featured projects
	projects, err := h.projectService.GetAllProjects(r.Context())
	if err != nil {
		h.logger.Error("Failed to get projects", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse template
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		h.logger.Error("Failed to parse template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":    "Portfolio - Home",
		"Projects": projects,
	}

	// Execute template
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		h.logger.Error("Failed to execute template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
