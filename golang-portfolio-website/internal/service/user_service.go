package service

import (
	"context"
	"portfolio/internal/model"
	"portfolio/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetProfile(ctx context.Context) (*model.User, error) {
	return s.repo.GetProfile(ctx)
}
