package main

import (
	"context"
	"github.com/marcustut/fyp/backend/config"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/migrate"
	"github.com/marcustut/fyp/backend/internal/infrastructure/datastore"
	"log"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("failed closing mysql client: %v", err)
		}
	}(client)
	createDBSchema(client)
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
