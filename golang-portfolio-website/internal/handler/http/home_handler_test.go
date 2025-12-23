package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.uber.org/zap"

	"portfolio/internal/model"
	"portfolio/internal/service"
)

type mockProjSvcForHome struct{}

func (m *mockProjSvcForHome) CreateProject(ctx context.Context, req interface{}) (interface{}, error) {
	return nil, nil
}
func (m *mockProjSvcForHome) GetProjectByID(ctx context.Context, id string) (interface{}, error) {
	return nil, nil
}
func (m *mockProjSvcForHome) GetAllProjects(ctx context.Context) ([]*model.Project, error) {
	return []*model.Project{{ID: "p1", Title: "Proj"}}, nil
}
func (m *mockProjSvcForHome) GetFeaturedProjects(ctx context.Context) ([]*model.Project, error) {
	return nil, nil
}

type mockUserSvcForHome struct{}

func (m *mockUserSvcForHome) GetProfile(ctx context.Context) (*model.User, error) { return nil, nil }

type mockUserRepoForHome struct{}

func (m *mockUserRepoForHome) GetByID(ctx context.Context, id string) (*model.User, error) {
	return nil, nil
}
func (m *mockUserRepoForHome) GetProfile(ctx context.Context) (*model.User, error) {
	return &model.User{ID: "u1", Name: "User"}, nil
}
func (m *mockUserRepoForHome) Update(ctx context.Context, user *model.User) error { return nil }

func TestHome_Renders(t *testing.T) {
	svc := service.NewProjectService(&mockRepoList{})
	userSvc := service.NewUserService(&mockUserRepoForHome{})
	h := NewHomeHandler(svc, userSvc, zap.NewExample())

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	h.Home(rr, r)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Proj") {
		t.Fatalf("expected body to contain project title")
	}
}
