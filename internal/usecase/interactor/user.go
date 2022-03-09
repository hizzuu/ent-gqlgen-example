package interactor

import (
	"context"

	"github.com/hizzuu/plate-backend/internal/domain"
	"github.com/hizzuu/plate-backend/internal/usecase/repository"
)

type user struct {
	userRepo  repository.User
	imageRepo repository.Image
	txRepo    repository.Transaction
	upRepo    repository.Uploader
}

type User interface {
	Get(ctx context.Context, id int) (*domain.User, error)
	GetByUID(ctx context.Context, uid string) (*domain.User, error)
	List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.UserWhereInput) (*domain.UserConnection, error)
	Create(ctx context.Context, input domain.CreateUserInput) (*domain.User, error)
	Update(ctx context.Context, input domain.UpdateUserInput) (*domain.User, error)
}

func NewUserInteractor(
	userRepo repository.User,
	imageRepo repository.Image,
	txRepo repository.Transaction,
	upRepo repository.Uploader,
) *user {
	return &user{
		userRepo:  userRepo,
		imageRepo: imageRepo,
		txRepo:    txRepo,
		upRepo:    upRepo,
	}
}

func (u *user) Get(ctx context.Context, id int) (*domain.User, error) {
	return u.userRepo.Get(ctx, id)
}

func (u *user) GetByUID(ctx context.Context, uid string) (*domain.User, error) {
	return u.userRepo.GetByUID(ctx, uid)
}

func (u *user) List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.UserWhereInput) (*domain.UserConnection, error) {
	return u.userRepo.List(ctx, after, first, before, last, where)
}

func (u *user) Create(ctx context.Context, input domain.CreateUserInput) (*domain.User, error) {
	return u.userRepo.Create(ctx, input)
}

func (u *user) Update(ctx context.Context, input domain.UpdateUserInput) (*domain.User, error) {
	user, err := u.txRepo.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		upImg := &domain.Image{File: input.AvatarImage.File}
		if err := resizeImage(upImg); err != nil {
			return nil, err
		}

		img, err := u.imageRepo.Create(ctx, domain.CreateImageInput{Name: upImg.Name})
		if err != nil {
			return nil, err
		}

		input.AvatarID = &img.ID
		user, err := u.userRepo.Update(ctx, input)
		if err != nil {
			return nil, err
		}

		if err := u.upRepo.Upload(ctx, upImg); err != nil {
			return nil, err
		}

		return user, nil
	})
	if err != nil {
		return nil, err
	}

	return user.(*domain.User), nil
}
