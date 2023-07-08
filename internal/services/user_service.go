package services

import (
	"github.com/devhijazi/go-users-api/pkg/database/models"
	"github.com/devhijazi/go-users-api/pkg/database/repositories"
	"github.com/devhijazi/go-users-api/pkg/errors"
)

type UserService interface {
	Create(user *models.User) (*models.User, *errors.Error)
	ListAll() []*models.User
	GetById(id string) (*models.User, *errors.Error)
	Update(id string, user *models.User) (*models.User, *errors.Error)
	// Delete(id string, user *models.User) (*models.User, *errors.Error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(
	userRepository repositories.UserRepository,
) UserService {
	return &userService{userRepository}
}

func (us *userService) Create(user *models.User) (*models.User, *errors.Error) {
	hasUserWithEmail, _ := us.userRepository.FindByEmail(user.Email)

	if hasUserWithEmail != nil {
		return nil, errors.UserNotFoundError()
	}

	userCreated := us.userRepository.Create(user)

	return userCreated, nil
}

func (us *userService) ListAll() []*models.User {
	userListAll := us.userRepository.FindAll()

	return userListAll
}

func (us *userService) GetById(id string) (*models.User, *errors.Error) {
	userGetById, err := us.userRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return userGetById, nil
}

func (us *userService) Update(id string, user *models.User) (*models.User, *errors.Error) {
	userUpdated, err := us.userRepository.Save(id, user)

	if err != nil {
		return nil, err
	}

	return userUpdated, nil

}

// func (us *userService) Delete(id string, user *models.User) (*models.User, *errors.Error) {
// 	userDeleted, err := us.userRepository.Delete(id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return userDeleted, nil

// }
