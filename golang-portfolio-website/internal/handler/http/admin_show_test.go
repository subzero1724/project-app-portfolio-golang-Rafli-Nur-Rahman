package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go.uber.org/zap"

	"portfolio/internal/model"
	"portfolio/internal/service"
)

type mockProjectRepoForAdmin struct{}

func (m *mockProjectRepoForAdmin) Create(ctx context.Context, p *model.Project) error {
	p.ID = "p1"
	return nil
}
func (m *mockProjectRepoForAdmin) GetByID(ctx context.Context, id string) (*model.Project, error) {
	return nil, nil
}
func (m *mockProjectRepoForAdmin) GetAll(ctx context.Context) ([]*model.Project, error) {
	return []*model.Project{{ID: "p1", Title: "Proj"}}, nil
}
func (m *mockProjectRepoForAdmin) GetFeatured(ctx context.Context) ([]*model.Project, error) {
	return nil, nil
}
func (m *mockProjectRepoForAdmin) Update(ctx context.Context, project *model.Project) error {
	return nil
}
func (m *mockProjectRepoForAdmin) Delete(ctx context.Context, id string) error { return nil }

type mockSkillRepoForAdmin struct{}

func (m *mockSkillRepoForAdmin) Create(ctx context.Context, s *model.Skill) error {
	s.ID = "s1"
	return nil
}
func (m *mockSkillRepoForAdmin) GetAll(ctx context.Context) ([]*model.Skill, error) {
	return []*model.Skill{{ID: "s1", Name: "Go"}}, nil
}

type mockExpRepoForAdmin struct{}

func (m *mockExpRepoForAdmin) Create(ctx context.Context, e *model.Experience) error {
	e.ID = "e1"
	return nil
}
func (m *mockExpRepoForAdmin) GetAll(ctx context.Context) ([]*model.Experience, error) {
	return []*model.Experience{{ID: "e1", Company: "C", Position: "P"}}, nil
}

type mockContactRepoForAdmin struct{}

func (m *mockContactRepoForAdmin) Create(ctx context.Context, c *model.Contact) error {
	c.ID = "m1"
	return nil
}
func (m *mockContactRepoForAdmin) GetByID(ctx context.Context, id string) (*model.Contact, error) {
	return nil, nil
}
func (m *mockContactRepoForAdmin) GetAll(ctx context.Context) ([]*model.Contact, error) {
	return []*model.Contact{{ID: "m1", Name: "A"}}, nil
}
func (m *mockContactRepoForAdmin) UpdateStatus(ctx context.Context, id string, status string) error {
	return nil
}
func (m *mockContactRepoForAdmin) Delete(ctx context.Context, id string) error { return nil }

func TestShowAdmin_Renders(t *testing.T) {
	// Ensure tests run from project root so templates can be found
	if _, err := os.Stat("templates/layout.html"); err != nil {
		// try moving up until found
		cur, _ := os.Getwd()
		found := false
		for i := 0; i < 5; i++ {
			cur = filepath.Join(cur, "..")
			if _, err := os.Stat(filepath.Join(cur, "templates/layout.html")); err == nil {
				os.Chdir(cur)
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("templates not found")
		}
	}
	projSvc := service.NewProjectService(&mockProjectRepoForAdmin{})
	skillSvc := service.NewSkillService(&mockSkillRepoForAdmin{})
	expSvc := service.NewExperienceService(&mockExpRepoForAdmin{})
	contactSvc := service.NewContactService(&mockContactRepoForAdmin{})

	h := NewAdminHandler(projSvc, skillSvc, expSvc, contactSvc, zap.NewExample())

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/admin", nil)
	h.ShowAdmin(rr, r)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Proj") {
		t.Fatalf("expected body to contain project title")
	}
}
