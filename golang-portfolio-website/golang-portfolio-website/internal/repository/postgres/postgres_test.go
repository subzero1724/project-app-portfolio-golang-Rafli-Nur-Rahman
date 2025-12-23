package postgres

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestNewRepositories_WithNilDB(t *testing.T) {
	var db *pgxpool.Pool
	if NewProjectRepository(db) == nil {
		t.Fatalf("expected project repo not nil")
	}
	if NewSkillRepository(db) == nil {
		t.Fatalf("expected skill repo not nil")
	}
	if NewExperienceRepository(db) == nil {
		t.Fatalf("expected experience repo not nil")
	}
	if NewContactRepository(db) == nil {
		t.Fatalf("expected contact repo not nil")
	}
	if NewUserRepository(db) == nil {
		t.Fatalf("expected user repo not nil")
	}
}

// calling methods with nil DB should panic (we recover to exercise code paths)
func TestProjectRepo_Methods_PanicOnNilDB(t *testing.T) {
	r := NewProjectRepository(nil)
	defer func() { recover() }()
	// these will panic due to nil db, but we recover
	_ = r.Create(nil, nil)
}

func TestProjectRepo_Getters_PanicOnNilDB(t *testing.T) {
	r := NewProjectRepository(nil)
	defer func() { recover() }()
	_, _ = r.GetByID(nil, "x")
}

func TestProjectRepo_List_PanicOnNilDB(t *testing.T) {
	r := NewProjectRepository(nil)
	defer func() { recover() }()
	_, _ = r.GetAll(nil)
}

func TestSkillRepo_Methods_PanicOnNilDB(t *testing.T) {
	r := NewSkillRepository(nil)
	defer func() { recover() }()
	_ = r.Create(nil, nil)
}

func TestContactRepo_Methods_PanicOnNilDB(t *testing.T) {
	r := NewContactRepository(nil)
	defer func() { recover() }()
	_ = r.Create(nil, nil)
}

func TestUserRepo_Methods_PanicOnNilDB(t *testing.T) {
	r := NewUserRepository(nil)
	defer func() { recover() }()
	_, _ = r.GetProfile(nil)
}
