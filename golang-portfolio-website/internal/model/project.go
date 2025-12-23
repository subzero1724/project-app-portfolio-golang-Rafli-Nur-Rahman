package model

import "time"

type Project struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	ImageURL    *string   `json:"image_url" db:"image_url"`
	DemoURL     *string   `json:"demo_url" db:"demo_url"`
	GithubURL   *string   `json:"github_url" db:"github_url"`
	TechStack   []string  `json:"tech_stack" db:"tech_stack"`
	IsFeatured  bool      `json:"is_featured" db:"is_featured"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
