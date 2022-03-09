package repository

import (
	"context"

	"github.com/hizzuu/plate-backend/internal/domain"
)

type User interface {
	Get(ctx context.Context, id int) (*domain.User, error)
	GetByUID(ctx context.Context, uid string) (*domain.User, error)
	List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.UserWhereInput) (*domain.UserConnection, error)
	Create(ctx context.Context, input domain.CreateUserInput) (*domain.User, error)
	Update(ctx context.Context, input domain.UpdateUserInput) (*domain.User, error)
}
