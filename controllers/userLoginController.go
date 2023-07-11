package controllers

import (
	"net/http"

	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"github.com/Pedro-Cecilio/projeto-livro-go/repositories"
	"github.com/Pedro-Cecilio/projeto-livro-go/services"
	"github.com/gin-gonic/gin"
)

func UserLoginController(c *gin.Context) {
	var loginData models.UserLogin
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userRepository := repositories.UserRepositoryImpl{}
	result, user := userRepository.FindUserByEmail(loginData)
	if !result {
		c.JSON(http.StatusNotFound, "Email não cadastrado")
		return
	}
	validPassword := services.CompareEncriptedPassword(loginData.Password, user.Password)
	if !validPassword{
		c.JSON(http.StatusNotFound, "Dados de usuário não encontrados")
		return
	}
	
	validTokenRepository := repositories.ValidTokenRepositoryImpl{}
	valid, validToken := validTokenRepository.CheckIfItIsValidById(user.ID)
	if !valid {
		jwt := services.GenerateJwt(user.ID)
		validTokenRepository.Create(user.ID, jwt)
		c.Header("Authorization", "Bearer "+ jwt)
	}else{
		c.Header("Authorization", "Bearer "+validToken.Jwt)
	}

	
	c.Status(http.StatusNoContent)
}
