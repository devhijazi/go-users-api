package controllers

import (
	"net/http"

	"github.com/devhijazi/go-users-api/internal/services"
	"github.com/devhijazi/go-users-api/pkg/database/repositories"
	"github.com/devhijazi/go-users-api/pkg/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sessionController struct {
	db *gorm.DB
}

func NewSessionController(db *gorm.DB) *sessionController {
	return &sessionController{db}
}

type RefreshSessionBody struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshSessionResponse struct {
	Session services.SessionRefreshReturnData `json:"session"`
}

// @Summary	Recuperar entidade
// @Description Recuperar sessão da entidade.
// @Tags sessions
// @Accept json
// @Produce json
// @Param refresh_token body string true "Refresh token"
// @Success 200 {object} RefreshSessionResponse
// @Failure 400 {object} errors.Error
// @Router /sessions/login/refresh [post]
func (sc *sessionController) RefreshSession(ctx *gin.Context) {
	var body *RefreshSessionBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(errors.ValidationError().GetStatus(), errors.ValidationError())

		return
	}

	userRepository := repositories.NewUserRepository(sc.db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(sc.db)

	sessionService := services.NewSessionService(userRepository, refreshTokenRepository)

	refreshData, err := sessionService.SessionRefresh(&services.SessionRefreshData{
		RefreshToken: body.RefreshToken,
	})

	if err != nil {
		ctx.JSON(err.GetStatus(), err.ToObject())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"session": refreshData})
}

type SessionUserLoginBody struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Stay     bool   `form:"stay"`
}

type SessionUserLoginResponse struct {
	Session services.SessionUserLoginReturnData `json:"session"`
}

// @Summary	Criar sessão do usuário
// @Description Criar sessão do usuário.
// @Tags sessions
// @Accept json
// @Produce json
// @Param email body string true "E-mail"
// @Param password body string true "Password"
// @Param stay body string true "Salvar sessão"
// @Success 200 {object} SessionUserLoginResponse
// @Failure 400 {object} errors.Error
// @Router /sessions/login/user [post]
func (sc *sessionController) SessionUserLogin(ctx *gin.Context) {
	var body *SessionUserLoginBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(errors.ValidationError().GetStatus(), errors.ValidationError())

		return
	}

	userRepository := repositories.NewUserRepository(sc.db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(sc.db)

	sessionService := services.NewSessionService(userRepository, refreshTokenRepository)

	loginData, err := sessionService.SessionUserLogin(&services.SessionUserLoginData{
		Email:    body.Email,
		Password: body.Password,
		Stay:     body.Stay,
	})

	if err != nil {
		ctx.JSON(err.GetStatus(), err.ToObject())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"session": loginData})
}
