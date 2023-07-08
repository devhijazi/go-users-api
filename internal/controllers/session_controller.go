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

type SessionUserLoginBody struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type SessionUserLoginResponse struct {
	Session services.SessionUserLoginReturnData `json:"session"`
}

func (sc *sessionController) SessionUserLogin(ctx *gin.Context) {
	var body *SessionUserLoginBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(errors.ValidationError().GetStatus(), errors.ValidationError())

		return
	}

	userRepository := repositories.NewUserRepository(sc.db)

	sessionService := services.NewSessionService(userRepository)

	loginData, err := sessionService.SessionUserLogin(&services.SessionUserLoginData{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.JSON(err.GetStatus(), err.ToObject())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"session": loginData})
}
