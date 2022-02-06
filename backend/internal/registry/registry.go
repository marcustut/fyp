package registry

import (
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
)

// all the dependencies goes here
type registry struct {
	client *ent.Client
}

// Registry is a wrapper over controllers for injecting
// the same dependencies into different controllers.
type Registry interface {
	NewController() controller.Controller
}

// New construct a Registry with the given dependencies.
func New(client *ent.Client) Registry {
	return &registry{client}
}

func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		Slide: r.NewSlideController(),
	}
}
