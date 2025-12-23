package dto

import "time"

type ProjectResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	DemoURL     string    `json:"demo_url"`
	GithubURL   string    `json:"github_url"`
	TechStack   []string  `json:"tech_stack"`
	IsFeatured  bool      `json:"is_featured"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
