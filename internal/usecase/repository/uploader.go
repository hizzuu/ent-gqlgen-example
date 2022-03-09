package repository

import (
	"context"

	"github.com/hizzuu/plate-backend/internal/domain"
)

type Uploader interface {
	Upload(ctx context.Context, img *domain.Image) error
	Delete(stx context.Context, fileName string) error
}
