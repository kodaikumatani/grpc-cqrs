package app

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-cqrs/internal/app/recipe"
	"github.com/kodaikumatani/grpc-cqrs/internal/app/user"
)

var Set = wire.NewSet(
	NewRegistrar,
	recipe.Set,
	user.Set,
)
