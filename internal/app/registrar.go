package app

import (
	pb "github.com/kodaikumatani/grpc-cqrs/pkg/pb/recipe"
	"google.golang.org/grpc"
)

type Registrar struct {
	recipeHandler pb.RecipeServiceServer
}

func NewRegistrar(
	recipeHandler pb.RecipeServiceServer,
) *Registrar {
	return &Registrar{
		recipeHandler: recipeHandler,
	}
}

func (r *Registrar) Register(app *grpc.Server) *grpc.Server {
	pb.RegisterRecipeServiceServer(app, r.recipeHandler)

	return app
}
