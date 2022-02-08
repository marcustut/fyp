package registry

import (
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
)

// Registry is a wrapper over controllers for injecting
// the same dependencies into different controllers.
type Registry interface {
	NewController() controller.Controller
}
