package registry

import (
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"github.com/marcustut/fyp/backend/internal/adapter/repository"
	"github.com/marcustut/fyp/backend/internal/usecase/usecase"
)

// all the dependencies for slide-service goes here
type slideRegistry struct {
	client *ent.Client
}

// NewSlideRegistry construct a Registry with the given dependencies
func NewSlideRegistry(client *ent.Client) Registry {
	return &slideRegistry{client}
}

// NewController implements Registry for slideRegistry
func (sr *slideRegistry) NewController() controller.Controller {
	return controller.Controller{
		Slide: sr.NewSlideController(),
	}
}

// NewSlideController conforms to interface
func (sr *slideRegistry) NewSlideController() controller.Slide {
	repo := repository.NewSlideRepository(sr.client)
	u := usecase.NewSlideUsecase(repo)
	return controller.NewSlideController(u)
}
