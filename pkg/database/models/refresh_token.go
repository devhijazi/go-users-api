package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        uuid.UUID `gorm:"primary_key" form:"id" json:"id,omitempty"`
	SubjectID string    `form:"subject_id" json:"subject_id,omitempty"`
	Token     string    `form:"token" json:"token,omitempty"`
}

func (refreshToken *RefreshToken) BeforeCreate(*gorm.DB) error {
	refreshToken.ID = uuid.New()

	return nil
}
