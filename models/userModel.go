package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string       `gorm:"type:char(36);not null;primaryKey"`
	Name     string       `gorm:"type:varchar(191);not null" binding:"required" json:"name"`
	Email    string       `gorm:"unique;not null" binding:"required,email" json:"email"`
	Password CustomString `gorm:"type:varchar(191);not null" binding:"required,min=8,max=20" json:"password"`

	BookRead	[]BookRead	`gorm:"foreignKey:User_id"`
	BooksBeingRead	[]BooksBeingRead	`gorm:"foreignKey:User_id"`
	BooksToRead	[]BooksToRead	`gorm:"foreignKey:User_id"`
	ValidToken []ValidToken `gorm:"foreignKey:User_id"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}
