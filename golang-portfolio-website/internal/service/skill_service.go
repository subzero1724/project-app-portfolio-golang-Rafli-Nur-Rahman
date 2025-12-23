package service

import (
	"context"
	"portfolio/internal/model"
	"portfolio/internal/repository"
)

type SkillService struct {
	repo repository.SkillRepository
}

func NewSkillService(repo repository.SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) CreateSkill(ctx context.Context, skill *model.Skill) error {
	return s.repo.Create(ctx, skill)
}

func (s *SkillService) GetAllSkills(ctx context.Context) ([]*model.Skill, error) {
	return s.repo.GetAll(ctx)
}
