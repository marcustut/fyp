package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marcustut/fyp/backend/config"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"github.com/marcustut/fyp/backend/internal/infrastructure/datastore"
	"github.com/marcustut/fyp/backend/internal/infrastructure/graphql"
	"github.com/marcustut/fyp/backend/internal/infrastructure/router"
	"github.com/marcustut/fyp/backend/internal/registry"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	ctrl := newController(client)

	srv := graphql.NewServer(client, ctrl)
	r := router.New(srv)

	log.Printf("slide-service running on port %d\n", config.C.Services.Slide.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.C.Services.Slide.Port), r))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}
	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.NewRegistry(client)
	return r.NewController()
}
