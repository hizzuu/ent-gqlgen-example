package storage

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/hizzuu/plate-backend/conf"
	"github.com/hizzuu/plate-backend/internal/domain"
)

type clientDev struct{}

type client struct {
	client *storage.Client
}

type Client interface {
	Upload(ctx context.Context, img *domain.Image) error
	Delete(stx context.Context, fileName string) error
}

func New(ctx context.Context) (Client, error) {
	if conf.C.App.Env == "dev" {
		return &clientDev{}, nil
	}

	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &client{client: c}, nil
}

// prod
func (c *client) Upload(ctx context.Context, img *domain.Image) error {
	return nil
}

func (c *client) Delete(stx context.Context, fileName string) error {
	return nil
}

// dev
func (c *clientDev) Upload(ctx context.Context, img *domain.Image) error {
	return nil
}

func (c *clientDev) Delete(stx context.Context, fileName string) error {
	return nil
}
