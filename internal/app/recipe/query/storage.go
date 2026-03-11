package query

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	Get(ctx context.Context, id uuid.UUID) (*RecipeWithUser, error)
	ListByUserID(ctx context.Context, userID string, limit, offset int32) ([]*Recipe, error)
}
