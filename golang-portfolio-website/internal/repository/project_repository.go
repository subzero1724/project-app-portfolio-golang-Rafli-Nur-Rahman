package repository

import (
	"context"
	"portfolio/internal/model"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *model.Project) error
	GetByID(ctx context.Context, id string) (*model.Project, error)
	GetAll(ctx context.Context) ([]*model.Project, error)
	GetFeatured(ctx context.Context) ([]*model.Project, error)
	Update(ctx context.Context, project *model.Project) error
	Delete(ctx context.Context, id string) error
}
