package main

import (
	"log"

	"github.com/devhijazi/go-users-api/internal/app"
	"github.com/devhijazi/go-users-api/pkg/database"
	"github.com/joho/godotenv"
)

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
