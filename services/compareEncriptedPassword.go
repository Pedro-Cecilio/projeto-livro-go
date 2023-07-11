package services

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"golang.org/x/crypto/bcrypt"
)

func CompareEncriptedPassword(loginPassword, DbPassword models.CustomString) bool {
	err := bcrypt.CompareHashAndPassword([]byte(DbPassword), []byte(loginPassword))
	
	return err == nil
}