package middlewares

import (
	"github.com/devhijazi/go-users-api/internal/services"
	"github.com/devhijazi/go-users-api/pkg/config"
	"github.com/devhijazi/go-users-api/pkg/database/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IsUserLoggedMiddleware(db *gorm.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {

		userAuthenticatedId := ctx.MustGet(config.AuthenticatedEntityIdConfig).(string)

		userRepository := repositories.NewUserRepository(db)
		refreshTokenRepository := repositories.NewRefreshTokenRepository(db)

		sessionService := services.NewSessionService(userRepository, refreshTokenRepository)

		err := sessionService.SessionVerifyIfIsUser(userAuthenticatedId)

		if err != nil {
			ctx.JSON(err.GetStatus(), err.ToObject())
			ctx.Abort()

			return
		}

		ctx.Next()

	})
}
