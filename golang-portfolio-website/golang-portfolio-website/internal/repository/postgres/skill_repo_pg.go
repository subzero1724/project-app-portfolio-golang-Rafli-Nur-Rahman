package postgres

import (
	"context"
	"portfolio/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SkillRepositoryPG struct {
	db *pgxpool.Pool
}

func NewSkillRepository(db *pgxpool.Pool) *SkillRepositoryPG {
	return &SkillRepositoryPG{db: db}
}

func (r *SkillRepositoryPG) Create(ctx context.Context, skill *model.Skill) error {
	query := `INSERT INTO skills (name, level, category) VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.db.QueryRow(ctx, query, skill.Name, skill.Level, skill.Category).Scan(&skill.ID, &skill.CreatedAt)
}

func (r *SkillRepositoryPG) GetAll(ctx context.Context) ([]*model.Skill, error) {
	query := `SELECT id, name, level, category, created_at FROM skills ORDER BY name`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []*model.Skill
	for rows.Next() {
		s := &model.Skill{}
		if err := rows.Scan(&s.ID, &s.Name, &s.Level, &s.Category, &s.CreatedAt); err != nil {
			return nil, err
		}
		skills = append(skills, s)
	}
	return skills, nil
}
