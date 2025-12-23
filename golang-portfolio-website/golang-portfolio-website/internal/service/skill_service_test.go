package service

import (
	"context"
	"testing"

	"portfolio/internal/model"
)

type mockSkillRepo struct{}

func (m *mockSkillRepo) Create(ctx context.Context, s *model.Skill) error { s.ID = "s1"; return nil }
func (m *mockSkillRepo) GetAll(ctx context.Context) ([]*model.Skill, error) {
	return []*model.Skill{{ID: "s1", Name: "Go"}}, nil
}

func TestSkillService_CreateAndList(t *testing.T) {
	svc := NewSkillService(&mockSkillRepo{})

	if err := svc.CreateSkill(context.Background(), &model.Skill{Name: "Go"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	skills, err := svc.GetAllSkills(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(skills) != 1 || skills[0].Name != "Go" {
		t.Fatalf("unexpected skills: %v", skills)
	}
}
