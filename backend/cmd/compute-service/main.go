package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/marcustut/fyp/backend/config"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	"github.com/marcustut/fyp/backend/internal/infrastructure/cloud"
	"github.com/marcustut/fyp/backend/internal/infrastructure/datastore"
	"github.com/marcustut/fyp/backend/internal/infrastructure/router"
	"github.com/marcustut/fyp/backend/internal/registry"
	"github.com/robfig/cron/v3"
	"github.com/rs/cors"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	cfg := newAWSConfig()
	ctrl := newController(client)

	r := router.NewCompute(ctrl, cfg)
	cr := newCron()
	cr.Start()

	log.Printf("compute-service running on port %d\n", config.C.Services.Compute.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.C.Services.Compute.Port), cors.Default().Handler(r)))
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

func newCron() *cron.Cron {
	c := cron.New()
	_, err := c.AddFunc("@every 5h", func() {
		// clean instances every hour
		reqURL := fmt.Sprintf("http://localhost:%d/clean", config.C.Services.Compute.Port)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			log.Printf("unable to create HTTP request: %v\n", err)
		}
		req.Header.Set("accept", "application/json")

		// execute the request
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("unable to send HTTP request: %v\n", err)
		}
		defer res.Body.Close()

		var msg string
		if err = json.NewDecoder(res.Body).Decode(&msg); err != nil {
			log.Printf("unable to parse JSON response: %v\n", err)
		}

		log.Println(msg)
	})
	if err != nil {
		log.Fatalf("error adding cron: %v", err)
	}
	return c
}
