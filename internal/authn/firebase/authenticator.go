package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/authn"
)

type verifier struct {
	client *auth.Client
}

func NewVerifier(ctx context.Context) (authn.Verifier, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &verifier{client: client}, nil
}

func (c *verifier) VerifyIDToken(ctx context.Context, idToken string) (context.Context, error) {
	token, err := c.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, authn.UIDKey{}, token.UID)
	ctx = context.WithValue(ctx, authn.ClaimsKey{}, token.Claims)
	return ctx, nil
}
