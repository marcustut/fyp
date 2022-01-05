package main

import (
	"fmt"
	"github.com/marcustut/fyp/backend/config"
	"github.com/marcustut/fyp/backend/ent"
	elk "github.com/marcustut/fyp/backend/ent/http"
	"github.com/marcustut/fyp/backend/internal/infrastructure/datastore"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()

	log.Printf("slide-service running on port %d\n", config.C.Services.Slide.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.C.Services.Slide.Port), elk.NewHandler(client, zap.NewExample())))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}
	return client
}
