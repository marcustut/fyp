package registry

import (
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"github.com/marcustut/fyp/backend/internal/adapter/repository"
	"github.com/marcustut/fyp/backend/internal/usecase/usecase"
)

// NewSlideController conforms to interface
func (r *registry) NewSlideController() controller.Slide {
	repo := repository.NewSlideRepository(r.client)
	u := usecase.NewSlideUsecase(repo)
	return controller.NewSlideController(u)
}
