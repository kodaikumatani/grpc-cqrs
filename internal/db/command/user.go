package command

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/app/user/command"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/app/user/domain"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/db/gen"
	"github.com/oklog/ulid/v2"
)

type user struct {
	queries *gen.Queries
}

func NewUser(pool *pgxpool.Pool) command.Storage {
	return &user{queries: gen.New(pool)}
}

func (u *user) Create(ctx context.Context, usr *domain.User) error {
	id, err := ulid.Parse(usr.ID)
	if err != nil {
		return err
	}

	return u.queries.CreateUser(ctx, gen.CreateUserParams{
		ID:        id,
		Name:      usr.Name,
		Email:     usr.Email,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	})
}
