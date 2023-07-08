package models

import (
	"time"

	"github.com/devhijazi/go-users-api/pkg/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key" form:"id" json:"id,omitempty"`
	FullName  string    `form:"full_name" json:"full_name,omitempty"`
	Email     string    `form:"email" json:"email,omitempty"`
	BirthDate string    `form:"birthdate" json:"birthdate,omitempty"`
	Phone     string    `form:"phone" json:"phone,omitempty"`
	Password  string    `form:"password" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" form:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" form:"updated_at" json:"updated_at,omitempty"`
}

func (user *User) BeforeCreate(*gorm.DB) error {
	passwordHash := helpers.GeneratePassword(user.Password)

	user.ID = uuid.New()
	user.Password = string(passwordHash)
	user.CreatedAt = time.Now().Local()

	return nil
}

func (user *User) BeforeUpdate(*gorm.DB) error {
	user.UpdatedAt = time.Now().Local()

	return nil
}
