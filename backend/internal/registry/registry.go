package registry

import "github.com/marcustut/fyp/backend/internal/adapter/controller"

// all the dependencies goes here
type registry struct {
}

// Registry is a wrapper over controllers for injecting
// the same dependencies into different controllers.
type Registry interface {
	NewController() controller.Controller
}

// New construct a Registry with the given dependencies.
func New() Registry {
	return &registry{}
}

func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		Slide: r.NewSlideController(),
	}
}
