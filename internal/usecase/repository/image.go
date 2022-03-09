package repository

import (
	"context"

	"github.com/hizzuu/plate-backend/internal/domain"
)

type Image interface {
	Create(ctx context.Context, input domain.CreateImageInput) (*domain.Image, error)
}
