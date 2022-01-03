package registry

import (
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"github.com/marcustut/fyp/backend/internal/adapter/repository"
	"github.com/marcustut/fyp/backend/internal/usecase"
)

func (r *registry) NewSlideController() controller.Slide {
	repo := repository.NewSlideRepository()
	u := usecase.NewSlideUsecase(repo)
	return controller.Slide(u)
}
