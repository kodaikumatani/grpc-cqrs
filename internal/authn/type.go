package authn

import "context"

type Verifier interface {
	VerifyIDToken(ctx context.Context, idToken string) (context.Context, error)
}

type UIDKey struct{}

type ClaimsKey struct{}
