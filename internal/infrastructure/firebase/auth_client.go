package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type authClient struct {
	client *auth.Client
}

type AuthClient interface {
	GetUser(ctx context.Context, uid string) (*auth.UserRecord, error)
	VerifyIDToken(ctx context.Context, token string) (*auth.Token, error)
	SetCustomClaims(ctx context.Context, uid string, id int) error
}

func NewAuthClient(ctx context.Context, app *firebase.App) (*authClient, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &authClient{client: client}, nil
}

func (a *authClient) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	return a.client.GetUser(ctx, uid)
}

func (a *authClient) VerifyIDToken(ctx context.Context, token string) (*auth.Token, error) {
	return a.client.VerifyIDToken(ctx, token)
}

func (a *authClient) SetCustomClaims(ctx context.Context, uid string, id int) error {
	claims := map[string]interface{}{"userID": id}
	if err := a.client.SetCustomUserClaims(ctx, uid, claims); err != nil {
		return err
	}

	return nil
}
