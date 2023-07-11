package controllers

import (
	"net/http"

	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"github.com/Pedro-Cecilio/projeto-livro-go/repositories"
	"github.com/Pedro-Cecilio/projeto-livro-go/services"
	"github.com/gin-gonic/gin"
)

func CreateUserController(c *gin.Context) {

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.Password = services.EncriptPassword(newUser.Password)
	userRepositoriy := repositories.UserRepositoryImpl{}
	result := userRepositoriy.Create(newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "Usuário criado com sucesso"})

}
