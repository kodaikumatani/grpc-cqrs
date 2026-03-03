package query

import "time"

type RecipeWithUser struct {
	ID          string
	UserID      string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserName    string
	UserEmail   string
}
