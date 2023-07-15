package controllers

import (
	"net/http"

	"github.com/devhijazi/go-users-api/internal/services"
	"github.com/devhijazi/go-users-api/pkg/config"
	"github.com/devhijazi/go-users-api/pkg/database/models"
	"github.com/devhijazi/go-users-api/pkg/database/repositories"
	"github.com/devhijazi/go-users-api/pkg/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type meController struct {
	db *gorm.DB
}

func NewMeController(db *gorm.DB) *meController {
	return &meController{db}
}

type SessionGetEntityResponse struct {
	User models.User `json:"user"`
}

// @Summary Obter entidade logada
// @Description	Retorna dados da entidade logada.
// @Tags me
// @Accep json
// @Produce json
// @Success 200 {object} SessionGetEntityResponse
// @Failure 400	{object} errors.Error
// @Router /me [get]

func (mc *meController) SessionGetEntity(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	userRepository := repositories.NewUserRepository(mc.db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(mc.db)

	sessionService := services.NewSessionService(userRepository, refreshTokenRepository)

	userData, entityError := sessionService.SessionGetEntity(authorization)

	if entityError != nil {
		ctx.JSON(entityError.GetStatus(), entityError.ToObject())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userData})

}

type UserUpdateBody struct {
	FullName string `form:"full_name" json:"full_name"`
	Email    string `form:"email" json:"email"`
	Phone    string `form:"phone" json:"phone"`
}

// @Summary Atualizar usuário
// @Description	Atualiza dados do usuário.
// @Tags me
// @Accept json
// @Produce json
// @Param full_name body string true "Nome completo"
// @Param email body string true "E-mail"
// @Param phone body string true "Telefone"
// @Success 200 {object} UserUpdateResponse
// @Failure 400 {object} errors.Error
// @Router /me/user [patch]
func (mc *meController) Update(ctx *gin.Context) {
	userId := ctx.MustGet(config.AuthenticatedEntityIdConfig).(string)

	var body *UserUpdateBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(errors.ValidationError().GetStatus(), errors.ValidationError())

		return
	}

	userRepository := repositories.NewUserRepository(mc.db)

	userService := services.NewUserService(userRepository)

	userUpdated, err := userService.Update(userId, &models.User{
		FullName: body.FullName,
		Email:    body.Email,
		Phone:    body.Phone,
	})

	if err != nil {
		ctx.JSON(err.GetStatus(), err.ToObject())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userUpdated})

}
