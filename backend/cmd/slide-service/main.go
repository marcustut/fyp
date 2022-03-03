package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/marcustut/fyp/backend/config"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"github.com/marcustut/fyp/backend/internal/infrastructure/cloud"
	"github.com/marcustut/fyp/backend/internal/infrastructure/datastore"
	"github.com/marcustut/fyp/backend/internal/infrastructure/graphql"
	"github.com/marcustut/fyp/backend/internal/infrastructure/router"
	"github.com/marcustut/fyp/backend/internal/registry"
	"github.com/rs/cors"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	cfg := newAWSConfig()
	s3 := newS3Client(cfg)
	ctrl := newController(client)

	srv := graphql.NewServer(client, ctrl, s3)
	r := router.New(srv)

	log.Printf("slide-service running on port %d\n", config.C.Services.Slide.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.C.Services.Slide.Port), cors.Default().Handler(r)))
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

func newAWSConfig() aws.Config {
	cfg, err := cloud.NewAWSConfig()
	if err != nil {
		log.Fatalf("failed getting aws config: %v", err)
	}
	return *cfg
}

func newS3Client(cfg aws.Config) *s3.Client {
	client := cloud.NewS3Client(cfg)
	return client
}
