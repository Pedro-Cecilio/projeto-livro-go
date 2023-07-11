package repositories

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/database"
	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) Create(newUser models.User) *gorm.DB {
	database.CreateConnection()
	result := database.DB.Create(&newUser)
	return result
}

func (u *UserRepositoryImpl) FindUserByEmail(loginData models.UserLogin) (bool, models.User) {
	database.CreateConnection()
	var user models.User
	result := database.DB.Where("email = ?", loginData.Email).First(&user)
	if result.RowsAffected == 0 {
		return false, user
	}
	return true, user
}

func EnsuresUserRepositoryImplImplementsUserRepository() {
	var _ userRepository = &UserRepositoryImpl{}
}
