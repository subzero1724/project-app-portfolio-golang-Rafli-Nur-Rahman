package postgres

import (
	"context"
	"portfolio/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ExperienceRepositoryPG struct {
	db *pgxpool.Pool
}

func NewExperienceRepository(db *pgxpool.Pool) *ExperienceRepositoryPG {
	return &ExperienceRepositoryPG{db: db}
}

func (r *ExperienceRepositoryPG) Create(ctx context.Context, exp *model.Experience) error {
	query := `INSERT INTO experiences (company, position, description, start_date, end_date, is_current) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, created_at`
	return r.db.QueryRow(ctx, query, exp.Company, exp.Position, exp.Description, exp.StartDate, exp.EndDate, exp.IsCurrent).Scan(&exp.ID, &exp.CreatedAt)
}

func (r *ExperienceRepositoryPG) GetAll(ctx context.Context) ([]*model.Experience, error) {
	query := `SELECT id, company, position, description, start_date, end_date, is_current, created_at FROM experiences ORDER BY start_date DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exps []*model.Experience
	for rows.Next() {
		e := &model.Experience{}
		if err := rows.Scan(&e.ID, &e.Company, &e.Position, &e.Description, &e.StartDate, &e.EndDate, &e.IsCurrent, &e.CreatedAt); err != nil {
			return nil, err
		}
		exps = append(exps, e)
	}
	return exps, nil
}
