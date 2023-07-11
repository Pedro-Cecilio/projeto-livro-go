package services

import (
	"fmt"

	"github.com/Pedro-Cecilio/projeto-livro-go/env"
	"github.com/golang-jwt/jwt/v5"
)

func GetPayloadUserIdFromJwt(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error){
		return []byte(env.Env.KEY_JWT), nil
	})
	if err != nil {
		return "", fmt.Errorf("ERRO AO ANALISAR TOKEN")
	}

	clains, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("ERRO AO OBTER AS CLAINS DO TOKEN")
	}

	payload := clains["user_id"].(string)

	return payload, nil
}
