package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hizzuu/plate-backend/ent"
	"github.com/hizzuu/plate-backend/graph/generated"
	"github.com/hizzuu/plate-backend/graph/model"
	"github.com/hizzuu/plate-backend/internal/domain"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput) (*model.CreatePostPayload, error) {
	u := ctx.Value(model.CurrentUserCtxKey).(*domain.User)
	input.UserID = u.ID
	p, err := r.controller.Post.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.CreatePostPayload{Post: p}, nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*model.GetPostPayload, error) {
	p, err := r.controller.Post.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.GetPostPayload{Post: p}, nil
}

func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	return r.controller.Post.List(ctx, after, first, before, last, where)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
