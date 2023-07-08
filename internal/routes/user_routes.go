package routes

import (
	"github.com/devhijazi/go-users-api/internal/controllers"
	"github.com/devhijazi/go-users-api/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	userRoute := r.Group("/users")

	userRoute.POST("", userController.Create)
	userRoute.GET("", middlewares.AuthenticationMiddleware(db), userController.ListAll)
	userRoute.GET("/:id", middlewares.AuthenticationMiddleware(db), userController.GetById)
	userRoute.PATCH("/:id", middlewares.AuthenticationMiddleware(db), userController.Update)
	// userRoute.DELETE("/:id", middlewares.AuthenticationMiddleware(db), userController.Delete)

}
