package repository

import (
	"context"
	"errors"

	"github.com/hizzuu/plate-backend/ent"
	"github.com/hizzuu/plate-backend/internal/domain"
)

type postRepository struct {
	client dbClient
}

func NewPostRepository(dbClient dbClient) *postRepository {
	return &postRepository{
		client: dbClient,
	}
}

func (r *postRepository) Get(ctx context.Context, id int) (*domain.Post, error) {
	p, err := r.client.Client(ctx).Post.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *postRepository) List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.PostWhereInput) (*domain.PostConnection, error) {
	ps, err := r.client.Client(ctx).Post.Query().WithUser().Paginate(ctx, after, first, before, last, ent.WithPostFilter(where.Filter))
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (r *postRepository) Create(ctx context.Context, input domain.CreatePostInput) (*domain.Post, error) {
	p, err := r.client.Client(ctx).Post.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *postRepository) Delete(ctx context.Context, id int, userID int) error {
	p, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	u, err := p.QueryUser().Only(ctx)
	if err != nil {
		return err
	}

	if u.ID != userID {
		return domain.NewForbiddenError(errors.New("this post cannot be deleted"))
	}

	return r.client.Client(ctx).Post.DeleteOne(p).Exec(ctx)
}
