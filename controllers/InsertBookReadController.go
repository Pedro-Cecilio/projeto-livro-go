package controllers

import (
	"net/http"
	"strings"

	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"github.com/Pedro-Cecilio/projeto-livro-go/repositories"
	"github.com/Pedro-Cecilio/projeto-livro-go/services"
	"github.com/gin-gonic/gin"
)

func InsertBookReadController(c *gin.Context) {
	var bookRead models.BookRead
	err := c.ShouldBindJSON(&bookRead)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	authorization := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authorization, "Bearer ")
	userid, err := services.GetPayloadUserIdFromJwt(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	bookRead.User_id = userid
	var bookRepository repositories.BookRepositoryImpl
	ok, err := bookRepository.AddBookRead(bookRead)
	if !ok {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusCreated)

}
