package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	s3         *s3.Client
}

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client, controller controller.Controller, s3 *s3.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Directives: generated.DirectiveRoot{},
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
			s3:         s3,
		},
	})
}
