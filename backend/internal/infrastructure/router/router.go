package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

// New creates route endpoint
func New(srv *handler.Server) *mux.Router {
	r := mux.NewRouter()

	r.Handle("/query", srv).Methods("POST")
	r.Handle("/", playground.Handler("GraphQL", "/query")).Methods("GET")

	return r
}
