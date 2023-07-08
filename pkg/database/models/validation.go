package models

import (
	"github.com/devhijazi/go-users-api/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Validation struct {
	ID                      uuid.UUID `gorm:"primary_key" form:"id" json:"id,omitempty"`
	Type                    string    `form:"type" json:"type,omitempty"`
	Validation              string    `form:"validation" json:"validation,omitempty"`
	Subject                 string    `form:"subject" json:"subject,omitempty"`
	ExpirationTimeInMinutes int       `form:"expiration_time_in_minutes" json:"expiration_time_in_minutes,omitempty"`
	CreatedTimestamp        int       `form:"created_timestamp" json:"created_timestamp,omitempty"`
}

func (validation *Validation) BeforeCreate(*gorm.DB) error {
	validation.ID = uuid.New()
	validation.CreatedTimestamp = utils.TimeNowInTimestamp()

	return nil
}
