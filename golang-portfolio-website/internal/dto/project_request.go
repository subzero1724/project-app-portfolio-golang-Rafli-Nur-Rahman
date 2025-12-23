package dto

type CreateProjectRequest struct {
	Title       string   `json:"title" validate:"required,min=3,max=100"`
	Description string   `json:"description" validate:"required,min=10"`
	ImageURL    string   `json:"image_url" validate:"omitempty,url"`
	DemoURL     string   `json:"demo_url" validate:"omitempty,url"`
	GithubURL   string   `json:"github_url" validate:"omitempty,url"`
	TechStack   []string `json:"tech_stack"`
	IsFeatured  bool     `json:"is_featured"`
}

type UpdateProjectRequest struct {
	Title       string   `json:"title" validate:"omitempty,min=3,max=100"`
	Description string   `json:"description" validate:"omitempty,min=10"`
	ImageURL    string   `json:"image_url" validate:"omitempty,url"`
	DemoURL     string   `json:"demo_url" validate:"omitempty,url"`
	GithubURL   string   `json:"github_url" validate:"omitempty,url"`
	TechStack   []string `json:"tech_stack"`
	IsFeatured  *bool    `json:"is_featured"`
}
