package registry

import (
	"github.com/hizzuu/plate-backend/internal/interface/controller"
	"github.com/hizzuu/plate-backend/internal/interface/repository"
	"github.com/hizzuu/plate-backend/internal/usecase/interactor"
)

func (r *registry) NewUserController() controller.User {
	return controller.NewUserController(
		interactor.NewUserInteractor(
			repository.NewUserRepository(r.dbHandler),
			repository.NewImageRepository(r.dbHandler),
			r.dbHandler,
			r.strHandler,
		),
	)
}
