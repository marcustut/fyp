package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/marcustut/fyp/backend/config"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"github.com/marcustut/fyp/backend/internal/infrastructure/cloud"
	"github.com/marcustut/fyp/backend/internal/infrastructure/datastore"
	"github.com/marcustut/fyp/backend/internal/registry"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	_ = newAWSConfig()
	ctrl := newController(client)

	ic, err := ctrl.Instance.List(context.TODO(), nil, nil, nil, nil, &ent.InstanceWhereInput{
		CreatedAtNotIn: []time.Time{time.Now().Add(-60 * time.Minute), time.Now()},
	}, nil)
	if err != nil {
		log.Fatalf("error listing instances: %v\n", err)
	}
	fmt.Printf("%+v\n", ic)
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
