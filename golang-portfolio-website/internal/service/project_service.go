package service

import (
	"context"
	"portfolio/internal/dto"
	"portfolio/internal/model"
	"portfolio/internal/repository"
)

type ProjectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *dto.CreateProjectRequest) (*dto.ProjectResponse, error) {
	var imageURL *string
	if req.ImageURL != "" {
		v := req.ImageURL
		imageURL = &v
	}
	var demoURL *string
	if req.DemoURL != "" {
		v := req.DemoURL
		demoURL = &v
	}
	var githubURL *string
	if req.GithubURL != "" {
		v := req.GithubURL
		githubURL = &v
	}

	project := &model.Project{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    imageURL,
		DemoURL:     demoURL,
		GithubURL:   githubURL,
		TechStack:   req.TechStack,
		IsFeatured:  req.IsFeatured,
	}

	if err := s.repo.Create(ctx, project); err != nil {
		return nil, err
	}

	return s.toResponse(project), nil
}

func (s *ProjectService) GetProjectByID(ctx context.Context, id string) (*dto.ProjectResponse, error) {
	project, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toResponse(project), nil
}

func (s *ProjectService) GetAllProjects(ctx context.Context) ([]*dto.ProjectResponse, error) {
	projects, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*dto.ProjectResponse
	for _, project := range projects {
		responses = append(responses, s.toResponse(project))
	}
	return responses, nil
}

func (s *ProjectService) GetFeaturedProjects(ctx context.Context) ([]*dto.ProjectResponse, error) {
	projects, err := s.repo.GetFeatured(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*dto.ProjectResponse
	for _, project := range projects {
		responses = append(responses, s.toResponse(project))
	}
	return responses, nil
}

func (s *ProjectService) toResponse(project *model.Project) *dto.ProjectResponse {
	image := ""
	if project.ImageURL != nil {
		image = *project.ImageURL
	}
	demo := ""
	if project.DemoURL != nil {
		demo = *project.DemoURL
	}
	github := ""
	if project.GithubURL != nil {
		github = *project.GithubURL
	}

	return &dto.ProjectResponse{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		ImageURL:    image,
		DemoURL:     demo,
		GithubURL:   github,
		TechStack:   project.TechStack,
		IsFeatured:  project.IsFeatured,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}
