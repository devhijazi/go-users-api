package routes

import (
	"github.com/devhijazi/go-users-api/internal/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SessionRoutes(r *gin.Engine, db *gorm.DB) {
	sessionController := controllers.NewSessionController(db)

	sessionRoute := r.Group("/sessions")

	sessionLoginRoute := sessionRoute.Group("/login")

	sessionLoginRoute.POST("/user", sessionController.SessionUserLogin)
	// sessionLoginRoute.POST("/refresh", sessionController.RefreshSession)
}
