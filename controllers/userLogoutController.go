package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Pedro-Cecilio/projeto-livro-go/repositories"
	"github.com/Pedro-Cecilio/projeto-livro-go/services"
	"github.com/gin-gonic/gin"
)

func UserLogoutController(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authorization, "Bearer ")
	userId, err := services.GetPayloadUserIdFromJwt(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}
	validTokenRepository := repositories.ValidTokenRepositoryImpl{}
	fmt.Println(userId)
	ok := validTokenRepository.InvalidateToken(userId)
	if !ok{
		c.JSON(http.StatusInternalServerError, "Erro ao invalidar token")
		return
	}
	c.Status(http.StatusNoContent)
}
