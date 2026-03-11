package internal

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/app"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/db"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/encrypt"
)

var Set = wire.NewSet(
	app.Set,
	db.Set,
	encrypt.Set,
)
