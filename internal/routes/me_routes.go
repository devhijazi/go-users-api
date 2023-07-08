package routes

import (
	"github.com/devhijazi/go-users-api/internal/controllers"
	"github.com/devhijazi/go-users-api/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MeRoutes(r *gin.Engine, db *gorm.DB) {
	meController := controllers.NewMeController(db)

	meRoute := r.Group("/me", middlewares.AuthenticationMiddleware(db))

	meRoute.GET("", meController.SessionGetEntity)

	meUserRoute := meRoute.Group("/user", middlewares.IsUserLoggedMiddleware(db))

	meUserRoute.PATCH("", meController.Update)
}
