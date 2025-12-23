package repository

import (
	"context"
	"portfolio/internal/model"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetProfile(ctx context.Context) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
}
