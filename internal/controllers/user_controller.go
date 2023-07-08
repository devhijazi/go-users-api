package controllers

import (
	"net/http"

	"github.com/devhijazi/go-users-api/internal/services"
	"github.com/devhijazi/go-users-api/pkg/database/models"
	"github.com/devhijazi/go-users-api/pkg/database/repositories"
	"github.com/devhijazi/go-users-api/pkg/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{db}
}

type UserCreateBody struct {
	FullName  string `form:"full_name" json:"full_name"`
	Email     string `form:"email" json:"email"`
	BirthDate string `form:"birthdate" json:"birthdate"`
	Phone     string `form:"phone" json:"phone"`
	Password  string `form:"password" json:"password"`
}

type UserCreateResponse struct {
	User models.User `json:"user"`
}

func (uc *userController) Create(ctx *gin.Context) {
	var body *UserCreateBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(errors.ValidationError().GetStatus(), errors.ValidationError())

		return
	}

	userRepository := repositories.NewUserRepository(uc.db)
	userService := services.NewUserService(userRepository)

	userCreated, err := userService.Create(&models.User{
		FullName:  body.FullName,
		Email:     body.Email,
		BirthDate: body.BirthDate,
		Phone:     body.Phone,
		Password:  body.Password,
	})

	if err != nil {
		ctx.JSON(err.GetStatus(), err.ToObject())
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": userCreated})
}

type UserListResponse struct {
	Users []models.User `json:"users"`
}

func (uc *userController) ListAll(ctx *gin.Context) {
	userRepository := repositories.NewUserRepository(uc.db)

	userService := services.NewUserService(userRepository)

	usersListed := userService.ListAll()

	ctx.JSON(http.StatusOK, gin.H{"users": usersListed})

}

type userGetByIdResponse struct {
	User models.User `json:"user"`
}

func (uc *userController) GetById(ctx *gin.Context) {
	userID := ctx.Param("id")

	userRepository := repositories.NewUserRepository(uc.db)

	userService := services.NewUserService(userRepository)

	userGetted, err := userService.GetById(userID)

	if err != nil {
		ctx.JSON(err.GetStatus(), err.ToObject())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userGetted})

}

type UserUpdateResponse struct {
	User models.User `json:"user"`
}

func (uc *userController) Update(ctx *gin.Context) {
	userID := ctx.Param("id")

	var body *UserCreateBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(errors.ValidationError().GetStatus(), errors.ValidationError())

		return
	}

	userRepository := repositories.NewUserRepository(uc.db)
	userService := services.NewUserService(userRepository)

	userUpdated, err := userService.Update(userID, &models.User{
		FullName:  body.FullName,
		Email:     body.Email,
		BirthDate: body.BirthDate,
		Phone:     body.Phone,
		Password:  body.Password,
	})

	if err != nil {
		ctx.JSON(err.GetStatus(), err.ToObject())
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userUpdated})

}

// func (uc *userController) Delete(ctx *gin.Context) {
// 	userID := ctx.Param("id")

// 	userRepository := repositories.NewUserRepository(uc.db)

// 	userService := services.NewUserService(userRepository)

// 	userDeleted, err := userService.Delete(userID, &models.User{})

// 	if err != nil {
// 		ctx.JSON(err.GetStatus(), err.ToObject())

// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"user": userDeleted})

// }
