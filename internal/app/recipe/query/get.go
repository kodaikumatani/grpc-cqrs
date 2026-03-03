package query

import (
	"context"

	"github.com/google/uuid"
)

type Query struct {
	storage Storage
}

func NewQuery(
	storage Storage,
) *Query {
	return &Query{
		storage: storage,
	}
}

func (q *Query) Get(
	ctx context.Context,
	id uuid.UUID,
) (*RecipeWithUser, error) {
	result, err := q.storage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
