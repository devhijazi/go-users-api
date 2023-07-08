package app

import (
	"fmt"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/devhijazi/go-users-api/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	cors "github.com/itsjamie/gin-cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/devhijazi/go-users-api/docs"
)

func StartServer(database *gorm.DB) {
	log.Println(color.InYellow("[Server] Starting server..."))

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.RedirectTrailingSlash = true

	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, PATCH, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
	}))

	routes.SessionRoutes(router, database)
	routes.MeRoutes(router, database)
	routes.UserRoutes(router, database)

	serverUrl := os.Getenv("URL")
	serverPort := fmt.Sprintf(":%s", os.Getenv("PORT"))

	router.GET(
		"/docs/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.URL(fmt.Sprintf("%s/docs/doc.json", serverUrl)),
		),
	)

	log.Printf(color.InGreen("[Server] Api started on %s"), serverUrl)
	log.Printf(color.InBlue("[Swagger Docs] See the api documentation at %s/docs/index.html"), serverUrl)

	router.Run(serverPort)
}
