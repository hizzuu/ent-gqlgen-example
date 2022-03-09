package resolver

import (
	"github.com/hizzuu/plate-backend/internal/interface/controller"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	controller controller.Controller
}

func New(controller controller.Controller) *Resolver {
	return &Resolver{
		controller: controller,
	}
}
