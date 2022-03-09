package registry

import (
	"github.com/hizzuu/plate-backend/internal/infrastructure/db"
	"github.com/hizzuu/plate-backend/internal/infrastructure/storage"
	"github.com/hizzuu/plate-backend/internal/interface/controller"
)

type registry struct {
	dbHandler  db.Client
	strHandler storage.Client
}

type Registry interface {
	NewController() controller.Controller
}

func New(dbHandler db.Client, strHandler storage.Client) *registry {
	return &registry{
		dbHandler:  dbHandler,
		strHandler: strHandler,
	}
}

// NewController generates controllers
func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		User: r.NewUserController(),
		Post: r.NewPostController(),
	}
}
