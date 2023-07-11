package services

import (
	"fmt"

	"github.com/Pedro-Cecilio/projeto-livro-go/env"
	"github.com/Pedro-Cecilio/projeto-livro-go/repositories"
	"github.com/golang-jwt/jwt/v5"
)

func JwtValidate(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("MÉTODO DE ASSINATURA INVÁLIDO")
		}
		return []byte(env.Env.KEY_JWT), nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	payload := token.Claims.(jwt.MapClaims)
	userID := payload["user_id"].(string)
	validTokenRepository := repositories.ValidTokenRepositoryImpl{}
	valid, _ := validTokenRepository.CheckIfItIsValidById(userID)
	if !valid {
		return false, nil
	}
	return true, nil

}
