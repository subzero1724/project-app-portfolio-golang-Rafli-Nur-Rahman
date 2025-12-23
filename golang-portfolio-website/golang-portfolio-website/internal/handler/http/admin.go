package http

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"portfolio/internal/dto"
	"portfolio/internal/model"
	"portfolio/internal/service"

	"go.uber.org/zap"
)

type AdminHandler struct {
	logger            *zap.Logger
	projectService    *service.ProjectService
	skillService      *service.SkillService
	experienceService *service.ExperienceService
	contactService    *service.ContactService
}

func NewAdminHandler(projectService *service.ProjectService, skillService *service.SkillService, experienceService *service.ExperienceService, contactService *service.ContactService, logger *zap.Logger) *AdminHandler {
	return &AdminHandler{logger: logger, projectService: projectService, skillService: skillService, experienceService: experienceService, contactService: contactService}
}

func (h *AdminHandler) ShowAdmin(w http.ResponseWriter, r *http.Request) {
	projects, _ := h.projectService.GetAllProjects(r.Context())
	skills, _ := h.skillService.GetAllSkills(r.Context())
	exps, _ := h.experienceService.GetAllExperiences(r.Context())
	msgs, _ := h.contactService.GetAllContacts(r.Context())

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/admin.html")
	if err != nil {
		h.logger.Error("Failed to parse template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":      "Admin",
		"Projects":   projects,
		"Skills":     skills,
		"Experience": exps,
		"Messages":   msgs,
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		h.logger.Error("Failed to execute template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// CreateSkill handles form POST to add a skill.
func (h *AdminHandler) CreateSkill(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	name := r.PostFormValue("name")
	levelStr := r.PostFormValue("level")
	category := r.PostFormValue("category")
	level, _ := strconv.Atoi(levelStr)

	// Basic validation
	if name == "" || level < 0 || level > 100 {
		http.Error(w, "Invalid skill input", http.StatusBadRequest)
		return
	}

	skill := &model.Skill{Name: name, Level: level, Category: category}
	if err := h.skillService.CreateSkill(r.Context(), skill); err != nil {
		h.logger.Error("Failed to create skill", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// redirect back to admin
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// CreateExperience handles form POST to add an experience.
func (h *AdminHandler) CreateExperience(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	company := r.PostFormValue("company")
	position := r.PostFormValue("position")
	description := r.PostFormValue("description")
	start := r.PostFormValue("start_date")
	end := r.PostFormValue("end_date")
	isCurrent := r.PostFormValue("is_current") == "on"

	sd, err := time.Parse("2006-01-02", start)
	if err != nil {
		http.Error(w, "Invalid start date", http.StatusBadRequest)
		return
	}
	var ed *time.Time
	if end != "" {
		t, err := time.Parse("2006-01-02", end)
		if err == nil {
			ed = &t
		}
	}

	if company == "" || position == "" {
		http.Error(w, "Invalid experience input", http.StatusBadRequest)
		return
	}

	exp := &model.Experience{Company: company, Position: position, Description: description, StartDate: sd, EndDate: ed, IsCurrent: isCurrent}
	if err := h.experienceService.CreateExperience(r.Context(), exp); err != nil {
		h.logger.Error("Failed to create experience", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// CreateProjectFromAdmin accepts form submissions to create project (accepts form or JSON)
func (h *AdminHandler) CreateProjectFromAdmin(w http.ResponseWriter, r *http.Request) {
	// Accept form POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	req := &dto.CreateProjectRequest{
		Title:       r.PostFormValue("title"),
		Description: r.PostFormValue("description"),
		ImageURL:    r.PostFormValue("image_url"),
		DemoURL:     r.PostFormValue("demo_url"),
		GithubURL:   r.PostFormValue("github_url"),
		TechStack:   []string{},
		IsFeatured:  r.PostFormValue("is_featured") == "on",
	}
	// tech stack as comma separated
	if ts := r.PostFormValue("tech_stack"); ts != "" {
		// split by comma
		var arr []string
		for _, t := range splitAndTrim(ts) {
			if t != "" {
				arr = append(arr, t)
			}
		}
		req.TechStack = arr
	}

	_, err := h.projectService.CreateProject(r.Context(), req)
	if err != nil {
		h.logger.Error("Failed to create project", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// helper to support basic ajax creation with JSON
func (h *AdminHandler) CreateSkillAPI(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Level    int    `json:"level"`
		Category string `json:"category"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid", http.StatusBadRequest)
		return
	}
	skill := &model.Skill{Name: req.Name, Level: req.Level, Category: req.Category}
	if err := h.skillService.CreateSkill(r.Context(), skill); err != nil {
		h.logger.Error("skill create", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func splitAndTrim(s string) []string {
	var out []string
	for _, p := range strings.Split(s, ",") {
		out = append(out, strings.TrimSpace(p))
	}
	return out
}
