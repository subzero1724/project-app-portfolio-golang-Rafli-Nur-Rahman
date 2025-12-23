package service

import (
	"context"
	"testing"

	"portfolio/internal/model"
)

type mockUserRepo struct{}

func (m *mockUserRepo) GetByID(ctx context.Context, id string) (*model.User, error) { return nil, nil }
func (m *mockUserRepo) GetProfile(ctx context.Context) (*model.User, error) {
	return &model.User{ID: "u1", Name: "User"}, nil
}
func (m *mockUserRepo) Update(ctx context.Context, user *model.User) error { return nil }

func TestUserService_GetProfile(t *testing.T) {
	svc := NewUserService(&mockUserRepo{})
	u, err := svc.GetProfile(context.Background())
	if err != nil || u.Name != "User" {
		t.Fatalf("unexpected user: %v %v", u, err)
	}
}
