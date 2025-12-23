package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"portfolio/internal/model"
	"strings"
	"testing"

	"portfolio/internal/service"

	"go.uber.org/zap"
)

type mockSkillRepo struct{}

func (m *mockSkillRepo) Create(ctx context.Context, s *model.Skill) error   { s.ID = "s1"; return nil }
func (m *mockSkillRepo) GetAll(ctx context.Context) ([]*model.Skill, error) { return nil, nil }

type mockExpRepo struct{}

func (m *mockExpRepo) Create(ctx context.Context, e *model.Experience) error   { e.ID = "e1"; return nil }
func (m *mockExpRepo) GetAll(ctx context.Context) ([]*model.Experience, error) { return nil, nil }

type mockProjectRepoForCreate struct{}

func (m *mockProjectRepoForCreate) Create(ctx context.Context, p *model.Project) error {
	p.ID = "p1"
	return nil
}
func (m *mockProjectRepoForCreate) GetByID(ctx context.Context, id string) (*model.Project, error) {
	return nil, nil
}
func (m *mockProjectRepoForCreate) GetAll(ctx context.Context) ([]*model.Project, error) {
	return nil, nil
}
func (m *mockProjectRepoForCreate) GetFeatured(ctx context.Context) ([]*model.Project, error) {
	return nil, nil
}
func (m *mockProjectRepoForCreate) Update(ctx context.Context, project *model.Project) error {
	return nil
}
func (m *mockProjectRepoForCreate) Delete(ctx context.Context, id string) error { return nil }

// Test CreateSkill form validation
func TestAdmin_CreateSkill_BadInput(t *testing.T) {
	skillSvc := service.NewSkillService(&mockSkillRepo{})
	h := &AdminHandler{logger: zap.NewExample(), skillService: skillSvc}

	form := url.Values{}
	form.Set("name", "")
	form.Set("level", "-5")

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/admin/skills", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	h.CreateSkill(rr, r)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

// Test CreateExperience date validation
func TestAdmin_CreateExperience_BadDate(t *testing.T) {
	expSvc := service.NewExperienceService(&mockExpRepo{})
	h := &AdminHandler{logger: zap.NewExample(), experienceService: expSvc}

	form := url.Values{}
	form.Set("company", "Acme")
	form.Set("position", "Dev")
	form.Set("start_date", "baddate")

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/admin/experience", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	h.CreateExperience(rr, r)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

// Test CreateExperience successful form submission
func TestAdmin_CreateExperience_Success(t *testing.T) {
	expSvc := service.NewExperienceService(&mockExpRepo{})
	h := &AdminHandler{logger: zap.NewExample(), experienceService: expSvc}

	form := url.Values{}
	form.Set("company", "Acme")
	form.Set("position", "Dev")
	form.Set("start_date", "2020-01-01")
	form.Set("is_current", "on")

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/admin/experience", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	h.CreateExperience(rr, r)
	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected 303, got %d", rr.Code)
	}
}

// Test CreateSkill successful form submission
func TestAdmin_CreateSkill_Success(t *testing.T) {
	skillSvc := service.NewSkillService(&mockSkillRepo{})
	h := &AdminHandler{logger: zap.NewExample(), skillService: skillSvc}

	form := url.Values{}
	form.Set("name", "Go")
	form.Set("level", "80")
	form.Set("category", "lang")

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/admin/skills", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	h.CreateSkill(rr, r)
	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected 303, got %d", rr.Code)
	}
}

// Test CreateSkillAPI successful JSON submission
func TestAdmin_CreateSkillAPI_Success(t *testing.T) {
	skillSvc := service.NewSkillService(&mockSkillRepo{})
	h := &AdminHandler{logger: zap.NewExample(), skillService: skillSvc}

	body := `{"name":"Go","level":90,"category":"lang"}`
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/admin/skills/api", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")

	h.CreateSkillAPI(rr, r)
	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rr.Code)
	}
}

// Test CreateProjectFromAdmin method and success flow
func TestAdmin_CreateProjectFromAdmin_Success(t *testing.T) {
	projSvc := service.NewProjectService(&mockProjectRepoForCreate{})
	h := &AdminHandler{logger: zap.NewExample(), projectService: projSvc}

	form := url.Values{}
	form.Set("title", "MyProj")
	form.Set("description", "desc")
	form.Set("tech_stack", "go, postgres, html")

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/admin/projects", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	h.CreateProjectFromAdmin(rr, r)
	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected 303, got %d", rr.Code)
	}
}

// Test splitAndTrim helper
func Test_splitAndTrim(t *testing.T) {
	out := splitAndTrim("a, b, c ,, d  ")
	if len(out) != 5 {
		t.Fatalf("unexpected length: %d", len(out))
	}
	if out[0] != "a" || out[1] != "b" || out[2] != "c" || out[3] != "" || out[4] != "d" {
		t.Fatalf("unexpected contents: %v", out)
	}
}
