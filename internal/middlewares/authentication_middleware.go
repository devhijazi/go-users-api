package middlewares

import (
	"strings"

	"github.com/devhijazi/go-users-api/pkg/config"
	"github.com/devhijazi/go-users-api/pkg/database/repositories"
	"github.com/devhijazi/go-users-api/pkg/errors"
	"github.com/devhijazi/go-users-api/pkg/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthenticationMiddleware(database *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")

		if authorization == "" {
			ctx.JSON(errors.TokenError().GetStatus(), errors.TokenError())
			ctx.Abort()

			return
		}

		splitAuthorization := strings.Split(authorization, " ")
		token := splitAuthorization[1]

		claims, err := helpers.ValidateAndDecodeSessionToken(token)

		if err != nil {
			ctx.JSON(err.GetStatus(), err.ToObject())
			ctx.Abort()

			return
		}

		entityId := claims.ID

		userRepository := repositories.NewUserRepository(database)

		user, err := userRepository.FindById(entityId)

		if err != nil {
			ctx.JSON(errors.AuthenticationError().GetStatus(), errors.AuthenticationError())
			ctx.Abort()
			return
		}

		ctx.Set(config.AuthenticatedEntityIdConfig, user.ID.String())

	}
}
