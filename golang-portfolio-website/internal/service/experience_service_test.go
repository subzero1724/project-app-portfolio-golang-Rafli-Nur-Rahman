package service

import (
	"context"
	"testing"

	"portfolio/internal/model"
)

type mockExpRepo struct{}

func (m *mockExpRepo) Create(ctx context.Context, e *model.Experience) error { e.ID = "e1"; return nil }
func (m *mockExpRepo) GetAll(ctx context.Context) ([]*model.Experience, error) {
	return []*model.Experience{{ID: "e1", Company: "C"}}, nil
}

func TestExperienceService_CreateAndList(t *testing.T) {
	svc := NewExperienceService(&mockExpRepo{})

	if err := svc.CreateExperience(context.Background(), &model.Experience{Company: "C"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	exps, err := svc.GetAllExperiences(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(exps) != 1 || exps[0].Company != "C" {
		t.Fatalf("unexpected exps: %v", exps)
	}
}
