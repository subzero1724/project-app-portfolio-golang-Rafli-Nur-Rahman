package model

import "time"

type User struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"-" db:"password"`
	Bio         string    `json:"bio" db:"bio"`
	AvatarURL   string    `json:"avatar_url" db:"avatar_url"`
	GithubURL   string    `json:"github_url" db:"github_url"`
	LinkedinURL string    `json:"linkedin_url" db:"linkedin_url"`
	TwitterURL  string    `json:"twitter_url" db:"twitter_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
