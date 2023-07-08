package database

import (
	"fmt"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/devhijazi/go-users-api/pkg/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDatabase() {
	log.Println(color.InYellow("[Database] Creating database connection..."))

	var _err error

	db, _err = gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"postgres://%s:%s@%s:%s/%s",
				os.Getenv("PG_USER"),
				os.Getenv("PG_PASSWORD"),
				os.Getenv("PG_HOST"),
				os.Getenv("PG_PORT"),
				os.Getenv("PG_DATABASE"),
			),
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if _err != nil {
		log.Println(color.InRed("[Database] Failed to connect to database."))
		log.Fatalln(_err)
	}

	log.Println(color.InGreen("[Database] connected to database!"))

	db.AutoMigrate(
		&models.User{},
		&models.Validation{},
		&models.RefreshToken{},
	)
}

func GetDatabase() *gorm.DB {
	return db
}
