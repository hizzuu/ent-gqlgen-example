package controller

import (
	"context"

	"github.com/hizzuu/plate-backend/internal/domain"
	"github.com/hizzuu/plate-backend/internal/usecase/interactor"
)

type post struct {
	postInteractor interactor.Post
}

type Post interface {
	Get(ctx context.Context, id int) (*domain.Post, error)
	List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.PostWhereInput) (*domain.PostConnection, error)
	Create(ctx context.Context, input domain.CreatePostInput) (*domain.Post, error)
	Delete(ctx context.Context, id int, userID int) error
}

func NewPostController(postInteractor interactor.Post) *post {
	return &post{
		postInteractor: postInteractor,
	}
}

func (c *post) Get(ctx context.Context, id int) (*domain.Post, error) {
	return c.postInteractor.Get(ctx, id)
}

func (c *post) List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.PostWhereInput) (*domain.PostConnection, error) {
	return c.postInteractor.List(ctx, after, first, before, last, where)
}

func (c *post) Create(ctx context.Context, input domain.CreatePostInput) (*domain.Post, error) {
	return c.postInteractor.Create(ctx, input)
}

func (c *post) Delete(ctx context.Context, id int, userID int) error {
	return c.postInteractor.Delete(ctx, id, userID)
}
