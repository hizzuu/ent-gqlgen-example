package registry

import (
	"github.com/hizzuu/plate-backend/internal/interface/controller"
	"github.com/hizzuu/plate-backend/internal/interface/repository"
	"github.com/hizzuu/plate-backend/internal/usecase/interactor"
)

func (r *registry) NewPostController() controller.Post {
	return controller.NewPostController(
		interactor.NewPostInteractor(
			repository.NewPostRepository(r.dbHandler),
			repository.NewImageRepository(r.dbHandler),
			r.dbHandler,
			r.strHandler,
		),
	)
}
