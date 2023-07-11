package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ValidToken struct {
	ID      string `gorm:"type:char(36);not null;primaryKey"`
	User_id string	`gorm:"not null"`
	Jwt     string `gorm:"type:varchar(256);not null"`
	IsValid bool
}

func (u *ValidToken) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}