package services

import (
	"strings"

	"github.com/devhijazi/go-users-api/pkg/database/models"
	"github.com/devhijazi/go-users-api/pkg/database/repositories"
	"github.com/devhijazi/go-users-api/pkg/errors"
	"github.com/devhijazi/go-users-api/pkg/helpers"
)

type SessionService interface {
	SessionGetEntity(authorization string) (*models.User, *errors.Error)
	SessionVerifyIfIsUser(id string) *errors.Error
	// SessionCreateRefreshToken(userId string) string
	// SessionRefresh(data *SessionRefreshData) (*SessionRefreshReturnData, *errors.Error)

	SessionUserLogin(data *SessionUserLoginData) (*SessionUserLoginReturnData, *errors.Error)
}

type sessionService struct {
	userRepository repositories.UserRepository
}

func NewSessionService(
	userRepository repositories.UserRepository,
	// refreshTokenRepository repositories.refreshTokenRepository,
) SessionService {
	return &sessionService{
		userRepository,
		// refreshTokenRepository,
	}
}

func (ss *sessionService) SessionGetEntity(authorization string) (*models.User, *errors.Error) {
	splitAuthorization := strings.Split(authorization, " ")

	token := splitAuthorization[1]

	claims, err := helpers.ValidateAndDecodeSessionToken(token)

	if err != nil {
		return nil, err
	}
	entityId := claims.ID

	user, _ := ss.userRepository.FindById(entityId)

	return user, errors.UserNotFoundError()
}

func (ss *sessionService) SessionVerifyIfIsUser(userId string) *errors.Error {
	_, err := ss.userRepository.FindById(userId)

	if err != nil {
		return errors.ForbiddenError()
	}

	return nil
}

// Adicioner o refresh token

type SessionUserLoginData struct {
	Email    string
	Password string
}

type SessionUserLoginReturnData struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

func (ss *sessionService) SessionUserLogin(data *SessionUserLoginData) (*SessionUserLoginReturnData, *errors.Error) {
	user, err := ss.userRepository.FindByEmail(data.Email)

	if err != nil {
		return nil, err
	}

	if helpers.ComparePassword(user.Password, data.Password) != nil {
		return nil, errors.PasswordError()
	}

	token := helpers.GenerateSessionToken(user.ID.String())

	return &SessionUserLoginReturnData{
		User:  user,
		Token: token,
	}, nil
}
