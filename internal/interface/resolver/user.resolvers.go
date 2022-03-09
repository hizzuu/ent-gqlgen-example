package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hizzuu/plate-backend/ent"
	"github.com/hizzuu/plate-backend/graph/model"
	"github.com/hizzuu/plate-backend/internal/domain"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*model.CreateUserPayload, error) {
	uid := ctx.Value(model.UIDCtxKey).(string)
	input.UID = uid
	input.Role = domain.RoleGeneral
	u, err := r.controller.User.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.CreateUserPayload{User: u}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input ent.UpdateUserInput) (*model.UpdateUserPayload, error) {
	user := ctx.Value(model.CurrentUserCtxKey).(*domain.User)
	input.ID = user.ID
	u, err := r.controller.User.Update(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.UpdateUserPayload{User: u}, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.GetUserPayload, error) {
	u, err := r.controller.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.GetUserPayload{User: u}, nil
}
