package service

import (
	"context"
	"testing"
	"time"

	"portfolio/internal/dto"
	"portfolio/internal/model"
)

type mockProjectRepo struct{}

func (m *mockProjectRepo) Create(ctx context.Context, project *model.Project) error {
	project.ID = "test-id"
	now := time.Now()
	project.CreatedAt = now
	project.UpdatedAt = now
	return nil
}
func (m *mockProjectRepo) GetByID(ctx context.Context, id string) (*model.Project, error) {
	return nil, nil
}
func (m *mockProjectRepo) GetAll(ctx context.Context) ([]*model.Project, error)      { return nil, nil }
func (m *mockProjectRepo) GetFeatured(ctx context.Context) ([]*model.Project, error) { return nil, nil }
func (m *mockProjectRepo) Update(ctx context.Context, project *model.Project) error  { return nil }
func (m *mockProjectRepo) Delete(ctx context.Context, id string) error               { return nil }

type mockProjectRepoGetters struct{}

func (r *mockProjectRepoGetters) Create(ctx context.Context, project *model.Project) error {
	return nil
}
func (r *mockProjectRepoGetters) GetByID(ctx context.Context, id string) (*model.Project, error) {
	return &model.Project{ID: id, Title: "T"}, nil
}
func (r *mockProjectRepoGetters) GetAll(ctx context.Context) ([]*model.Project, error) {
	return []*model.Project{{ID: "p1", Title: "P1"}}, nil
}
func (r *mockProjectRepoGetters) GetFeatured(ctx context.Context) ([]*model.Project, error) {
	return []*model.Project{{ID: "f1", Title: "F1"}}, nil
}
func (r *mockProjectRepoGetters) Update(ctx context.Context, project *model.Project) error {
	return nil
}
func (r *mockProjectRepoGetters) Delete(ctx context.Context, id string) error { return nil }

func TestCreateProject_WithURLs_SetsPointersAndReturnsResponse(t *testing.T) {
	repo := &mockProjectRepo{}
	svc := NewProjectService(repo)

	req := &dto.CreateProjectRequest{
		Title:       "My Project",
		Description: "A nice project",
		ImageURL:    "https://example.com/img.png",
		DemoURL:     "https://example.com/demo",
		GithubURL:   "https://github.com/repo",
		TechStack:   []string{"Go", "Postgres"},
		IsFeatured:  true,
	}

	resp, err := svc.CreateProject(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.ID != "test-id" {
		t.Fatalf("expected id set by repo, got %s", resp.ID)
	}
	if resp.ImageURL == "" || resp.DemoURL == "" || resp.GithubURL == "" {
		t.Fatalf("expected urls to be preserved in response")
	}
}

func TestProjectService_Getters(t *testing.T) {
	// repo that returns specific data
	svc := NewProjectService(&mockProjectRepoGetters{})

	p, err := svc.GetProjectByID(context.Background(), "p1")
	if err != nil || p.Title != "T" {
		t.Fatalf("unexpected project: %v %v", p, err)
	}

	all, err := svc.GetAllProjects(context.Background())
	if err != nil || len(all) != 1 || all[0].Title != "P1" {
		t.Fatalf("unexpected all: %v %v", all, err)
	}

	f, err := svc.GetFeaturedProjects(context.Background())
	if err != nil || len(f) != 1 || f[0].Title != "F1" {
		t.Fatalf("unexpected featured: %v %v", f, err)
	}
}
