package service

import (
	"context"
	"portfolio/internal/model"
	"portfolio/internal/repository"
)

type ExperienceService struct {
	repo repository.ExperienceRepository
}

func NewExperienceService(repo repository.ExperienceRepository) *ExperienceService {
	return &ExperienceService{repo: repo}
}

func (s *ExperienceService) CreateExperience(ctx context.Context, exp *model.Experience) error {
	return s.repo.Create(ctx, exp)
}

func (s *ExperienceService) GetAllExperiences(ctx context.Context) ([]*model.Experience, error) {
	return s.repo.GetAll(ctx)
}
