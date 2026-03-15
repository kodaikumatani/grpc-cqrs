package interceptor

import (
	"context"
	"errors"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/authn"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidAuthToken = errors.New("invalid authorization token")
)

func AuthUnaryInterceptor(verifier authn.Verifier) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		tokenString, err := auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, ErrInvalidAuthToken.Error())
		}
		if tokenString == "" {
			return nil, status.Error(codes.Unauthenticated, ErrInvalidAuthToken.Error())
		}

		ctx, err = verifier.VerifyIDToken(ctx, tokenString)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, ErrInvalidAuthToken.Error())
		}

		return handler(ctx, req)
	}
}
