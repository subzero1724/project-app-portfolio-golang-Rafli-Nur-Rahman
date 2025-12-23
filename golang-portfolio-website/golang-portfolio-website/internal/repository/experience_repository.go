package repository

import (
	"context"
	"portfolio/internal/model"
)

type ExperienceRepository interface {
	Create(ctx context.Context, exp *model.Experience) error
	GetAll(ctx context.Context) ([]*model.Experience, error)
}
