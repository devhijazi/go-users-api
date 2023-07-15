package repositories

import (
	"github.com/devhijazi/go-users-api/pkg/database/models"
	"github.com/devhijazi/go-users-api/pkg/errors"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(refreshToken *models.RefreshToken)
	FindByToken(token string) (*models.RefreshToken, *errors.Error)
	Remove(token string) *errors.Error
}

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{db}
}

func (rr *refreshTokenRepository) Create(refreshToken *models.RefreshToken) {
	rr.db.Create(&refreshToken)
}

func (rr *refreshTokenRepository) FindByToken(token string) (*models.RefreshToken, *errors.Error) {
	var refreshToken *models.RefreshToken

	if rr.db.Where("token = ?", token).First(&refreshToken).Error != nil {
		return nil, errors.UserNotFoundError()
	}

	return refreshToken, nil
}

func (rr *refreshTokenRepository) Remove(token string) *errors.Error {
	refreshToken, err := rr.FindByToken(token)

	if err != nil {
		return err
	}

	rr.db.Delete(&refreshToken)

	return nil
}
