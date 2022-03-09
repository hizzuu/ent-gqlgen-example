package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/hizzuu/plate-backend/conf"
)

func New(ctx context.Context) (*firebase.App, error) {
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON([]byte(conf.C.Api.FirebaseKeyJson)))
	if err != nil {
		return nil, err
	}

	return app, nil
}
