package controllers

import (
	"github.com/devhijazi/go-users-api/pkg/database/models"
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

func (mc *meController) SessionGetEntity(ctx *gin.Context) {
	// TODO
}
func (mc *meController) Update(ctx *gin.Context) {
	// TODO
}
