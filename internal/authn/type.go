package authn

import (
	"context"
	"errors"
)

var ErrUnauthenticated = errors.New("user is not authenticated")

type Verifier interface {
	VerifyIDToken(ctx context.Context, idToken string) (context.Context, error)
}

type UIDKey struct{}

type ClaimsKey struct{}
