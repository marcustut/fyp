package registry

import (
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"github.com/marcustut/fyp/backend/internal/adapter/repository"
	"github.com/marcustut/fyp/backend/internal/usecase/usecase"
)

// Registry is a wrapper over controllers for injecting
// the same dependencies into different controllers.
type Registry interface {
	NewController() controller.Controller
}

// all the dependencies for auth-service goes here
type registry struct {
	client *ent.Client
}

// NewRegistry construct a Registry with the given dependencies
func NewRegistry(client *ent.Client) Registry {
	return &registry{client}
}

// NewController implements Registry for registry
func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		User:     r.NewUserController(),
		Slide:    r.NewSlideController(),
		Instance: r.NewInstanceController(),
	}
}

// NewUserController conforms to interface
func (r *registry) NewUserController() controller.User {
	repo := repository.NewUserRepository(r.client)
	u := usecase.NewUserUsecase(repo)
	return controller.NewUserController(u)
}

// NewSlideController conforms to interface
func (r *registry) NewSlideController() controller.Slide {
	repo := repository.NewSlideRepository(r.client)
	u := usecase.NewSlideUsecase(repo)
	return controller.NewSlideController(u)
}

// NewInstanceController conforms to interface
func (r *registry) NewInstanceController() controller.Instance {
	repo := repository.NewInstanceRepository(r.client)
	u := usecase.NewInstanceUsecase(repo)
	return controller.NewInstanceController(u)
}
