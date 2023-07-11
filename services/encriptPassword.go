package services

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"golang.org/x/crypto/bcrypt"
)

func EncriptPassword(password models.CustomString)models.CustomString{
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Erro ao encriptar senha")
	}
	return models.CustomString(hash)
}
