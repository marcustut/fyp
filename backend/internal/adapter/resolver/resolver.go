package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/graph/generated"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is a context struct
type Resolver struct {
	client     *ent.Client
	controller controller.Controller
}

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Directives: generated.DirectiveRoot{},
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
		},
	})
}
