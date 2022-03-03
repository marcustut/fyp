package graphql

import (
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"github.com/marcustut/fyp/backend/internal/adapter/resolver"
)

// NewServer generates graphql server
func NewServer(client *ent.Client, controller controller.Controller, s3 *s3.Client) *handler.Server {
	srv := handler.NewDefaultServer(resolver.NewSchema(client, controller, s3))
	srv.Use(entgql.Transactioner{TxOpener: client})

	return srv
}
