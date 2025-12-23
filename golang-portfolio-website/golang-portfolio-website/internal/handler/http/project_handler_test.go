package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"portfolio/internal/dto"
	"portfolio/internal/model"
	"portfolio/internal/service"
)

type mockRepoForHandler struct{}

func (m *mockRepoForHandler) Create(_ context.Context, project *model.Project) error {
	project.ID = "h-test-id"
	return nil
}
func (m *mockRepoForHandler) GetByID(_ context.Context, id string) (*model.Project, error) {
	return nil, nil
}
func (m *mockRepoForHandler) GetAll(_ context.Context) ([]*model.Project, error) { return nil, nil }
func (m *mockRepoForHandler) GetFeatured(_ context.Context) ([]*model.Project, error) {
	return nil, nil
}
func (m *mockRepoForHandler) Update(_ context.Context, project *model.Project) error { return nil }
func (m *mockRepoForHandler) Delete(_ context.Context, id string) error              { return nil }

// helpers for specific tests
type mockRepoList struct{}

func (m *mockRepoList) Create(ctx context.Context, project *model.Project) error { return nil }
func (m *mockRepoList) GetByID(ctx context.Context, id string) (*model.Project, error) {
	return nil, nil
}
func (m *mockRepoList) GetAll(ctx context.Context) ([]*model.Project, error) {
	return []*model.Project{{ID: "p1", Title: "Proj"}}, nil
}
func (m *mockRepoList) GetFeatured(ctx context.Context) ([]*model.Project, error) { return nil, nil }
func (m *mockRepoList) Update(ctx context.Context, project *model.Project) error  { return nil }
func (m *mockRepoList) Delete(ctx context.Context, id string) error               { return nil }

type mockRepoGetNotFound struct{}

func (m *mockRepoGetNotFound) Create(ctx context.Context, project *model.Project) error { return nil }
func (m *mockRepoGetNotFound) GetByID(ctx context.Context, id string) (*model.Project, error) {
	return nil, fmt.Errorf("not found")
}
func (m *mockRepoGetNotFound) GetAll(ctx context.Context) ([]*model.Project, error) { return nil, nil }
func (m *mockRepoGetNotFound) GetFeatured(ctx context.Context) ([]*model.Project, error) {
	return nil, nil
}
func (m *mockRepoGetNotFound) Update(ctx context.Context, project *model.Project) error { return nil }
func (m *mockRepoGetNotFound) Delete(ctx context.Context, id string) error              { return nil }

type mockRepoGetOK struct{}

func (m *mockRepoGetOK) Create(ctx context.Context, project *model.Project) error { return nil }
func (m *mockRepoGetOK) GetByID(ctx context.Context, id string) (*model.Project, error) {
	return &model.Project{ID: id, Title: "Proj"}, nil
}
func (m *mockRepoGetOK) GetAll(ctx context.Context) ([]*model.Project, error)      { return nil, nil }
func (m *mockRepoGetOK) GetFeatured(ctx context.Context) ([]*model.Project, error) { return nil, nil }
func (m *mockRepoGetOK) Update(ctx context.Context, project *model.Project) error  { return nil }
func (m *mockRepoGetOK) Delete(ctx context.Context, id string) error               { return nil }

// Test validation rejects missing title
func TestCreateProject_ValidationFails(t *testing.T) {
	svc := service.NewProjectService(&mockRepoForHandler{})
	h := NewProjectHandler(svc, zap.NewExample())

	req := dto.CreateProjectRequest{Description: "short", Title: ""}
	b, _ := json.Marshal(req)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/projects", bytes.NewReader(b))
	h.CreateProject(rr, r)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

// Test successful creation returns 201
func TestCreateProject_Success(t *testing.T) {
	svc := service.NewProjectService(&mockRepoForHandler{})
	h := NewProjectHandler(svc, zap.NewExample())

	req := dto.CreateProjectRequest{Description: "A valid description long enough", Title: "A Title"}
	b, _ := json.Marshal(req)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/projects", bytes.NewReader(b))
	h.CreateProject(rr, r)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}
	var resp dto.ProjectResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.Title != req.Title {
		t.Fatalf("expected title %s, got %s", req.Title, resp.Title)
	}
}

// Test ListProjects HTML rendering
func TestListProjects_HTML(t *testing.T) {
	// ensure templates available (like admin test does)
	// try to find templates by moving up if necessary
	// (reuse admin_show_test approach)
	// Note: keep this simple; if templates can't be found, skip test
	// so CI isn't brittle.
	rr := httptest.NewRecorder()

	svc := service.NewProjectService(&mockRepoList{})
	h := NewProjectHandler(svc, zap.NewExample())

	r := httptest.NewRequest(http.MethodGet, "/projects", nil)
	h.ListProjects(rr, r)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Proj") {
		t.Fatalf("expected body to contain project title")
	}
}

// Test API list returns JSON
func TestListProjects_API(t *testing.T) {
	svc := service.NewProjectService(&mockRepoList{})
	h := NewProjectHandler(svc, zap.NewExample())

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/projects", nil)
	h.ListProjects(rr, r)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Proj") {
		t.Fatalf("expected JSON to contain project title")
	}
}

// Test GetProject not found
func TestGetProject_NotFound(t *testing.T) {
	svc := service.NewProjectService(&mockRepoGetNotFound{})
	h := NewProjectHandler(svc, zap.NewExample())

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/projects/p1", nil)
	// chi.URLParam isn't set unless using router; set it in context
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "p1")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	h.GetProject(rr, r)
	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rr.Code)
	}
}

// Test GetProject success
func TestGetProject_Success(t *testing.T) {
	svc := service.NewProjectService(&mockRepoGetOK{})
	h := NewProjectHandler(svc, zap.NewExample())

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/projects/p1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "p1")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	h.GetProject(rr, r)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Proj") {
		t.Fatalf("expected body to contain project title")
	}
}
