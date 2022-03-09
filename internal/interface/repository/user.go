package repository

import (
	"context"

	"github.com/hizzuu/plate-backend/ent"
	"github.com/hizzuu/plate-backend/ent/user"
	"github.com/hizzuu/plate-backend/internal/domain"
)

type userRepository struct {
	client dbClient
}

func NewUserRepository(client dbClient) *userRepository {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) Get(ctx context.Context, id int) (*domain.User, error) {
	u, err := r.client.Client(ctx).User.Get(ctx, id)
	if err != nil {
		if ent.IsNotSingular(err) || ent.IsNotFound(err) {
			return nil, domain.NewNotFoundError(err, map[string]interface{}{
				"id": id,
			})
		}

		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetByUID(ctx context.Context, uid string) (*domain.User, error) {
	u, err := r.client.Client(ctx).User.Query().Where(user.UIDEQ(uid)).Only(ctx)
	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, domain.NewNotFoundError(err, map[string]interface{}{
				"uid": uid,
			})
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	return u, nil
}

func (r *userRepository) List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.UserWhereInput) (*domain.UserConnection, error) {
	us, err := r.client.Client(ctx).User.Query().Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter))
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (r *userRepository) Create(ctx context.Context, input domain.CreateUserInput) (*domain.User, error) {
	u, err := r.client.Client(ctx).User.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) Update(ctx context.Context, input domain.UpdateUserInput) (*domain.User, error) {
	u, err := r.client.Client(ctx).User.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
