package interactor

import (
	"context"

	"github.com/hizzuu/plate-backend/internal/domain"
	"github.com/hizzuu/plate-backend/internal/usecase/repository"
)

type post struct {
	postRepo  repository.Post
	imageRepo repository.Image
	txRepo    repository.Transaction
	upRepo    repository.Uploader
}

type Post interface {
	Get(ctx context.Context, id int) (*domain.Post, error)
	List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.PostWhereInput) (*domain.PostConnection, error)
	Create(ctx context.Context, input domain.CreatePostInput) (*domain.Post, error)
	Delete(ctx context.Context, id int, userID int) error
}

func NewPostInteractor(
	postRepo repository.Post,
	imageRepo repository.Image,
	txRepo repository.Transaction,
	upRepo repository.Uploader,
) *post {
	return &post{
		postRepo:  postRepo,
		imageRepo: imageRepo,
		txRepo:    txRepo,
		upRepo:    upRepo,
	}
}

func (p *post) Get(ctx context.Context, id int) (*domain.Post, error) {
	return p.postRepo.Get(ctx, id)
}

func (p *post) List(ctx context.Context, after *domain.Cursor, first *int, before *domain.Cursor, last *int, where *domain.PostWhereInput) (*domain.PostConnection, error) {
	return p.postRepo.List(ctx, after, first, before, last, where)
}

func (p *post) Create(ctx context.Context, input domain.CreatePostInput) (*domain.Post, error) {
	post, err := p.txRepo.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		upImg := &domain.Image{File: input.PhotoImage.File}
		if err := resizeImage(upImg); err != nil {
			return nil, err
		}

		img, err := p.imageRepo.Create(ctx, domain.CreateImageInput{Name: upImg.Name})
		if err != nil {
			return nil, err
		}

		input.PhotoID = img.ID
		post, err := p.postRepo.Create(ctx, input)
		if err != nil {
			return nil, err
		}

		if err := p.upRepo.Upload(ctx, upImg); err != nil {
			return nil, err
		}

		return post, nil
	})
	if err != nil {
		return nil, err
	}

	return post.(*domain.Post), nil
}

func (p *post) Delete(ctx context.Context, id int, userID int) error {
	return p.postRepo.Delete(ctx, id, userID)
}
