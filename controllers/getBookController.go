package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"github.com/gin-gonic/gin"
)

func GetBookController(c *gin.Context){
	id := c.Param("id")
	var url string
	if id != ""{
		url = "https://www.googleapis.com/books/v1/volumes/" + id
	}else{
		url = "https://www.googleapis.com/books/v1/volumes/" 
	}
	response, err := http.Get(url) 
	if err != nil {
		c.JSON(response.StatusCode, err.Error())
		return 
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(response.StatusCode, err.Error())
		return
	}
	var book models.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		c.JSON(response.StatusCode, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)

	

}