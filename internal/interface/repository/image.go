package repository

import (
	"context"

	"github.com/hizzuu/plate-backend/internal/domain"
)

type imageRepository struct {
	dbClient dbClient
}

func NewImageRepository(dbClient dbClient) *imageRepository {
	return &imageRepository{
		dbClient: dbClient,
	}
}

func (r *imageRepository) Create(ctx context.Context, input domain.CreateImageInput) (*domain.Image, error) {
	i, err := r.dbClient.Client(ctx).Image.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return i, nil
}
