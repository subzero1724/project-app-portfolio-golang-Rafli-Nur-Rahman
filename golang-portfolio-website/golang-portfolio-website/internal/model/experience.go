package model

import "time"

type Experience struct {
	ID          string     `json:"id" db:"id"`
	Company     string     `json:"company" db:"company"`
	Position    string     `json:"position" db:"position"`
	Description string     `json:"description" db:"description"`
	StartDate   time.Time  `json:"start_date" db:"start_date"`
	EndDate     *time.Time `json:"end_date" db:"end_date"`
	IsCurrent   bool       `json:"is_current" db:"is_current"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}
