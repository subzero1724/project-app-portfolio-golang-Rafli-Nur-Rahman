package repository

import (
	"context"
	"portfolio/internal/model"
)

type SkillRepository interface {
	Create(ctx context.Context, skill *model.Skill) error
	GetAll(ctx context.Context) ([]*model.Skill, error)
}
