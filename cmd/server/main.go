package main

import (
	"log"

	"github.com/devhijazi/go-users-api/internal/app"
	"github.com/devhijazi/go-users-api/pkg/database"
	"github.com/joho/godotenv"
)

// @title Users API - GoLang
// @version 0.0.1
// @description Complete user API with PostgreSQL database and Kafka for event streaming.
// @termsOfService	https://gabrielhijazi.dev/terms

// @contact.name	Gabriel Hijazi
// @contact.url http://gabrielhijazi.dev
// @contact.email hijazi.root@gmail.com"

// @license.name Copyright (c) Gabriel Hijazi Fraga Santos
// @license.url http://gabrielhijazi.dev

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log.SetFlags(2 | 3)

	godotenv.Load()

	database.ConnectDatabase()

	// kafka.ConnectKafka()

	app.StartServer(
		database.GetDatabase(),
		// Kafka.GetProducer()
	)
}
