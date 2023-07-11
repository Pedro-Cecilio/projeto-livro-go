package repositories

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/database"
	"github.com/Pedro-Cecilio/projeto-livro-go/models"
)

type ValidTokenRepositoryImpl struct {
}

func (vt *ValidTokenRepositoryImpl) Create(User_id, jwt string) {
	database.CreateConnection()
	validToken := models.ValidToken{
		User_id: User_id,
		Jwt:     jwt,
		IsValid: true,
	}
	database.DB.Create(&validToken)
}

func (vt *ValidTokenRepositoryImpl) CheckIfItIsValidById(User_id string) (bool, models.ValidToken) {
	database.CreateConnection()
	var validToken models.ValidToken
	result := database.DB.Where("user_id = ? AND is_valid = ?", User_id, true).First(&validToken)
	return result.RowsAffected != 0, validToken

}

func (vt *ValidTokenRepositoryImpl) InvalidateToken(User_id string) bool {
	database.CreateConnection()
	var validToken models.ValidToken

	result := database.DB.Model(&validToken).Where("user_id = ? AND is_valid = ?", User_id, true).Update("is_valid", false)
	return result.RowsAffected != 0

}

func EnsuresValidTokenRepositoryImplImplementsValidTokenRepository() {
	var _ validToken = &ValidTokenRepositoryImpl{}
}
