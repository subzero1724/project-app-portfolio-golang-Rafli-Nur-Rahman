package postgres

import (
	"context"
	"portfolio/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepositoryPG struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepositoryPG {
	return &ProjectRepositoryPG{db: db}
}

func (r *ProjectRepositoryPG) Create(ctx context.Context, project *model.Project) error {
	query := `
		INSERT INTO projects (title, description, image_url, demo_url, github_url, tech_stack, is_featured)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRow(ctx, query,
		project.Title,
		project.Description,
		project.ImageURL,
		project.DemoURL,
		project.GithubURL,
		project.TechStack,
		project.IsFeatured,
	).Scan(&project.ID, &project.CreatedAt, &project.UpdatedAt)
}

func (r *ProjectRepositoryPG) GetByID(ctx context.Context, id string) (*model.Project, error) {
	query := `
		SELECT id, title, description, image_url, demo_url, github_url, tech_stack, is_featured, created_at, updated_at
		FROM projects WHERE id = $1
	`
	project := &model.Project{}
	var imageURL *string
	var demoURL *string
	var githubURL *string

	err := r.db.QueryRow(ctx, query, id).Scan(
		&project.ID,
		&project.Title,
		&project.Description,
		&imageURL,
		&demoURL,
		&githubURL,
		&project.TechStack,
		&project.IsFeatured,
		&project.CreatedAt,
		&project.UpdatedAt,
	)

	project.ImageURL = imageURL
	project.DemoURL = demoURL
	project.GithubURL = githubURL
	return project, err
}

func (r *ProjectRepositoryPG) GetAll(ctx context.Context) ([]*model.Project, error) {
	query := `
		SELECT id, title, description, image_url, demo_url, github_url, tech_stack, is_featured, created_at, updated_at
		FROM projects ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*model.Project
	for rows.Next() {
		project := &model.Project{}
		var imageURL *string
		var demoURL *string
		var githubURL *string

		err := rows.Scan(
			&project.ID,
			&project.Title,
			&project.Description,
			&imageURL,
			&demoURL,
			&githubURL,
			&project.TechStack,
			&project.IsFeatured,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		project.ImageURL = imageURL
		project.DemoURL = demoURL
		project.GithubURL = githubURL
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (r *ProjectRepositoryPG) GetFeatured(ctx context.Context) ([]*model.Project, error) {
	query := `
		SELECT id, title, description, image_url, demo_url, github_url, tech_stack, is_featured, created_at, updated_at
		FROM projects WHERE is_featured = true ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*model.Project
	for rows.Next() {
		project := &model.Project{}
		var imageURL *string
		var demoURL *string
		var githubURL *string

		err := rows.Scan(
			&project.ID,
			&project.Title,
			&project.Description,
			&imageURL,
			&demoURL,
			&githubURL,
			&project.TechStack,
			&project.IsFeatured,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		project.ImageURL = imageURL
		project.DemoURL = demoURL
		project.GithubURL = githubURL
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (r *ProjectRepositoryPG) Update(ctx context.Context, project *model.Project) error {
	query := `
		UPDATE projects
		SET title = $1, description = $2, image_url = $3, demo_url = $4, 
		    github_url = $5, tech_stack = $6, is_featured = $7, updated_at = NOW()
		WHERE id = $8
	`
	_, err := r.db.Exec(ctx, query,
		project.Title,
		project.Description,
		project.ImageURL,
		project.DemoURL,
		project.GithubURL,
		project.TechStack,
		project.IsFeatured,
		project.ID,
	)
	return err
}

func (r *ProjectRepositoryPG) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM projects WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
