package middleware

import (
	"net/http"
	"strings"

	"github.com/Pedro-Cecilio/projeto-livro-go/services"
	"github.com/gin-gonic/gin"
)

func CheckIfAuthorizationIsValid() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authorization,"Bearer ")
		valid, err := services.JwtValidate(token)
		if !valid || err != nil{
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
