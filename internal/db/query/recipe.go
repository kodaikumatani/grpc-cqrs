package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/app/recipe/query"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/db/gen"
)

type recipe struct {
	queries *gen.Queries
}

func NewRecipe(pool *pgxpool.Pool) query.Storage {
	return &recipe{queries: gen.New(pool)}
}

func (r *recipe) ListByUserID(ctx context.Context, userID string, limit, offset int32) ([]*query.Recipe, error) {
	rows, err := r.queries.ListRecipesByUserID(ctx, gen.ListRecipesByUserIDParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*query.Recipe, 0, len(rows))
	for _, row := range rows {
		result = append(result, &query.Recipe{
			ID:          row.ID.String(),
			UserID:      row.UserID,
			Title:       row.Title,
			Description: row.Description,
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
		})
	}

	return result, nil
}

func (r *recipe) Get(ctx context.Context, id uuid.UUID) (*query.RecipeWithUser, error) {
	row, err := r.queries.GetRecipeWithUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &query.RecipeWithUser{
		ID:          row.ID.String(),
		UserID:      row.UserID,
		Title:       row.Title,
		Description: row.Description,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		UserName:    row.UserName,
		UserEmail:   row.UserEmail,
	}, nil
}
