package repositories

import (
	"github.com/devhijazi/go-users-api/pkg/database/models"
	"github.com/devhijazi/go-users-api/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) *models.User
	FindAll() []*models.User
	FindById(id string) (*models.User, *errors.Error)
	FindByEmail(email string) (*models.User, *errors.Error)
	Save(id string, user *models.User) (*models.User, *errors.Error)
	UpdatePassword(id string, password string) *errors.Error
	// Delete(id string) (*models.User, *errors.Error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Create(user *models.User) *models.User {
	ur.db.Create(&user)

	userCreated, _ := ur.FindById(user.ID.String())

	return userCreated
}

func (ur *userRepository) FindAll() []*models.User {
	var users []*models.User

	ur.db.Model(&models.User{}).Find(&users)

	return users
}

func (ur *userRepository) FindById(id string) (*models.User, *errors.Error) {
	var user *models.User

	if ur.db.Where("id = ?", id).First(&user).Error != nil {
		return nil, errors.UserNotFoundError()
	}

	return user, nil
}

func (ur *userRepository) FindByEmail(email string) (*models.User, *errors.Error) {
	var user *models.User

	if ur.db.Where("email = ?", email).First(&user).Error != nil {
		return nil, errors.UserNotFoundError()
	}

	return user, nil
}

func (ur *userRepository) Save(id string, user *models.User) (*models.User, *errors.Error) {
	userFounded, err := ur.FindById(id)

	if err != nil {
		return nil, err
	}

	userFounded.FullName = user.FullName
	userFounded.Email = user.Email
	userFounded.Phone = user.Phone

	ur.db.Save(&userFounded)

	return userFounded, nil
}

func (ur *userRepository) UpdatePassword(id string, password string) *errors.Error {
	userFounded, err := ur.FindById(id)

	if err != nil {
		return err
	}

	userFounded.Password = password
	ur.db.Save(&userFounded)

	return nil
}

func (ur *userRepository) Delete(id string) *errors.Error {
	userFounded, err := ur.FindById(id)

	if err != nil {
		return err
	}

	ur.db.Delete(&userFounded)

	return nil
}
