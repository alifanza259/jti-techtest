package main

import (
	"log"

	"github.com/alifanza259/jwt-techtest/api"
	"github.com/alifanza259/jwt-techtest/models"
	"github.com/alifanza259/jwt-techtest/storage"
	"github.com/alifanza259/jwt-techtest/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatalf("cannot connect to database: %s", err)
	}
	models.Migrate(db)

	server := api.NewServer(config, db)

	err = server.Start(config.Host)
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
