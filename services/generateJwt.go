package services

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/env"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(user_id string) string {
	clains := jwt.MapClaims{
		"user_id": user_id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
	signedToken, err := token.SignedString([]byte(env.Env.KEY_JWT))
	if err != nil {
		panic("Erro ao assinar token")
	}

	return signedToken
}
