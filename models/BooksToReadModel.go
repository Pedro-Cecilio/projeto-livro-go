package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BooksToRead struct {
	ID      string `gorm:"type:char(36);not null;primaryKey"`
	User_id string	`gorm:"not null"`
	Book_id	string  `gorm:"type:char(36);not null" json:"book_id"`
	gorm.Model
}

func (u *BooksToRead) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}
