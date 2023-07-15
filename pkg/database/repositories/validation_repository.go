package repositories

import (
	"github.com/devhijazi/go-users-api/pkg/database/models"
	"github.com/devhijazi/go-users-api/pkg/errors"
	"gorm.io/gorm"
)

type ValidationRepository interface {
	Create(validation *models.Validation) *models.Validation
	FindById(id string) (*models.Validation, *errors.Error)
	FindByTypeAndSubject(typee string, subject string) (*models.Validation, *errors.Error)
	FindByValidationAndType(validation string, typee string) (*models.Validation, *errors.Error)
	Remove(typee string, subject string) *errors.Error
}

type validationRepository struct {
	db *gorm.DB
}

func NewValidationRepository(db *gorm.DB) ValidationRepository {
	return &validationRepository{db}
}

func (vr *validationRepository) Create(validation *models.Validation) *models.Validation {
	vr.db.Create(&validation)

	validationCreated, _ := vr.FindById(validation.ID.String())

	return validationCreated
}

func (vr *validationRepository) FindById(id string) (*models.Validation, *errors.Error) {
	var validation *models.Validation

	if vr.db.Where("id = ?", id).First(&validation).Error != nil {
		return nil, errors.UserNotFoundError()
	}

	return validation, nil
}

func (vr *validationRepository) FindByTypeAndSubject(typee string, subject string) (*models.Validation, *errors.Error) {
	var validation *models.Validation

	if vr.db.Where("type = ? AND subject = ?", typee, subject).First(&validation).Error != nil {
		return nil, errors.UserNotFoundError()
	}

	return validation, nil
}

func (vr *validationRepository) FindByValidationAndType(validationn string, typee string) (*models.Validation, *errors.Error) {
	var validation *models.Validation

	if vr.db.Where("validation = ? AND type = ?", validationn, typee).First(&validation).Error != nil {
		return nil, errors.UserNotFoundError()
	}

	return validation, nil
}

func (vr *validationRepository) Remove(typee string, subject string) *errors.Error {
	validation, err := vr.FindByTypeAndSubject(typee, subject)

	if err != nil {
		return err
	}

	vr.db.Delete(&validation)

	return nil
}
