package http

import (
	"encoding/json"
	"html/template"
	"net/http"
	"portfolio/internal/dto"
	"portfolio/internal/service"
	"portfolio/internal/util"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type ProjectHandler struct {
	projectService *service.ProjectService
	logger         *zap.Logger
}

func NewProjectHandler(projectService *service.ProjectService, logger *zap.Logger) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
		logger:         logger,
	}
}

func (h *ProjectHandler) ListProjects(w http.ResponseWriter, r *http.Request) {
	// Check if this is an API request
	if r.Header.Get("Accept") == "application/json" || r.URL.Path == "/api/projects" {
		h.listProjectsAPI(w, r)
		return
	}

	// HTML response
	projects, err := h.projectService.GetAllProjects(r.Context())
	if err != nil {
		h.logger.Error("Failed to get projects", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/projects.html")
	if err != nil {
		h.logger.Error("Failed to parse template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":    "Portfolio - Projects",
		"Projects": projects,
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		h.logger.Error("Failed to execute template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *ProjectHandler) listProjectsAPI(w http.ResponseWriter, r *http.Request) {
	projects, err := h.projectService.GetAllProjects(r.Context())
	if err != nil {
		h.logger.Error("Failed to get projects", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	project, err := h.projectService.GetProjectByID(r.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get project", zap.Error(err))
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if err := util.ValidateStruct(&req); err != nil {
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	project, err := h.projectService.CreateProject(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to create project", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}
