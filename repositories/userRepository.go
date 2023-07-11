package repositories

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"gorm.io/gorm"
)

type userRepository interface {
	Create(newUser models.User) *gorm.DB
	FindUserByEmail(loginData models.UserLogin) (bool, models.User)
}

