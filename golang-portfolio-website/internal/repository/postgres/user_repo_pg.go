package postgres

import (
	"context"
	"portfolio/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryPG struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepositoryPG {
	return &UserRepositoryPG{db: db}
}

func (r *UserRepositoryPG) GetByID(ctx context.Context, id string) (*model.User, error) {
	query := `
		SELECT id, name, email, bio, avatar_url, github_url, linkedin_url, twitter_url, created_at, updated_at
		FROM users WHERE id = $1
	`
	user := &model.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Bio,
		&user.AvatarURL,
		&user.GithubURL,
		&user.LinkedinURL,
		&user.TwitterURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}

func (r *UserRepositoryPG) GetProfile(ctx context.Context) (*model.User, error) {
	query := `
		SELECT id, name, email, bio, avatar_url, github_url, linkedin_url, twitter_url, created_at, updated_at
		FROM users LIMIT 1
	`
	user := &model.User{}
	err := r.db.QueryRow(ctx, query).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Bio,
		&user.AvatarURL,
		&user.GithubURL,
		&user.LinkedinURL,
		&user.TwitterURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}

func (r *UserRepositoryPG) Update(ctx context.Context, user *model.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, bio = $3, avatar_url = $4,
			github_url = $5, linkedin_url = $6, twitter_url = $7, updated_at = NOW()
		WHERE id = $8
	`
	_, err := r.db.Exec(ctx, query,
		user.Name,
		user.Email,
		user.Bio,
		user.AvatarURL,
		user.GithubURL,
		user.LinkedinURL,
		user.TwitterURL,
		user.ID,
	)
	return err
}
