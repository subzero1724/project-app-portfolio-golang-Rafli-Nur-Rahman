package router

import (
	"portfolio/internal/handler/http"
	"portfolio/internal/repository/postgres"
	"portfolio/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func NewRouter(db *pgxpool.Pool, logger *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Static files
	// fileServer := http.FileServer(http.Dir("./static"))
	// r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Repositories
	projectRepo := postgres.NewProjectRepository(db)
	contactRepo := postgres.NewContactRepository(db)
	userRepo := postgres.NewUserRepository(db)
	skillRepo := postgres.NewSkillRepository(db)
	expRepo := postgres.NewExperienceRepository(db)

	// Services
	projectService := service.NewProjectService(projectRepo)
	contactService := service.NewContactService(contactRepo)
	userService := service.NewUserService(userRepo)
	skillService := service.NewSkillService(skillRepo)
	expService := service.NewExperienceService(expRepo)

	// Handlers
	healthHandler := http.NewHealthHandler()
	homeHandler := http.NewHomeHandler(projectService, userService, logger)
	projectHandler := http.NewProjectHandler(projectService, logger)
	contactHandler := http.NewContactHandler(contactService, logger)
	adminHandler := http.NewAdminHandler(projectService, skillService, expService, contactService, logger)

	// Routes
	r.Get("/health", healthHandler.Health)
	r.Get("/", homeHandler.Home)
	r.Get("/projects", projectHandler.ListProjects)
	r.Get("/contact", contactHandler.ShowContactForm)
	r.Get("/admin", adminHandler.ShowAdmin)

	// Admin form endpoints
	r.Post("/admin/skills", adminHandler.CreateSkill)
	r.Post("/admin/experience", adminHandler.CreateExperience)
	r.Post("/admin/projects", adminHandler.CreateProjectFromAdmin)

	// API Routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/projects", func(r chi.Router) {
			r.Get("/", projectHandler.ListProjects)
			r.Post("/", projectHandler.CreateProject)
			r.Get("/{id}", projectHandler.GetProject)
		})

		r.Route("/contact", func(r chi.Router) {
			r.Post("/", contactHandler.SubmitContact)
		})
	})

	return r
}
