package repository

import (
	"context"

	"github.com/hizzuu/plate-backend/internal/domain"
)

type Post interface {
	Get(ctx context.Context, id int) (*domain.Post, error)
	List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.PostWhereInput) (*domain.PostConnection, error)
	Create(ctx context.Context, input domain.CreatePostInput) (*domain.Post, error)
	Delete(ctx context.Context, id int, userID int) error
}
