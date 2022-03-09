package controller

import (
	"context"

	"github.com/hizzuu/plate-backend/internal/domain"
	"github.com/hizzuu/plate-backend/internal/usecase/interactor"
)

type user struct {
	userInteractor interactor.User
}

type User interface {
	Get(ctx context.Context, id int) (*domain.User, error)
	GetByUID(ctx context.Context, uid string) (*domain.User, error)
	List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.UserWhereInput) (*domain.UserConnection, error)
	Create(ctx context.Context, input domain.CreateUserInput) (*domain.User, error)
	Update(ctx context.Context, input domain.UpdateUserInput) (*domain.User, error)
}

func NewUserController(userInteractor interactor.User) *user {
	return &user{
		userInteractor: userInteractor,
	}
}

func (u *user) Get(ctx context.Context, id int) (*domain.User, error) {
	return u.userInteractor.Get(ctx, id)
}

func (u *user) GetByUID(ctx context.Context, uid string) (*domain.User, error) {
	return u.userInteractor.GetByUID(ctx, uid)
}

func (u *user) List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.UserWhereInput) (*domain.UserConnection, error) {
	return u.userInteractor.List(ctx, after, first, before, last, where)
}

func (u *user) Create(ctx context.Context, input domain.CreateUserInput) (*domain.User, error) {
	return u.userInteractor.Create(ctx, input)
}

func (u *user) Update(ctx context.Context, input domain.UpdateUserInput) (*domain.User, error) {
	return u.userInteractor.Update(ctx, input)
}
