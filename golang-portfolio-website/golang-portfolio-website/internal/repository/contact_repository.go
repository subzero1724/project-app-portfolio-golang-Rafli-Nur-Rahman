package repository

import (
	"context"
	"portfolio/internal/model"
)

type ContactRepository interface {
	Create(ctx context.Context, contact *model.Contact) error
	GetByID(ctx context.Context, id string) (*model.Contact, error)
	GetAll(ctx context.Context) ([]*model.Contact, error)
	UpdateStatus(ctx context.Context, id string, status string) error
	Delete(ctx context.Context, id string) error
}
