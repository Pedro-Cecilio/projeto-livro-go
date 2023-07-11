package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"github.com/gin-gonic/gin"
)

func SearchBooksController(c *gin.Context) {
	search := c.Query("search")
	var url string
	if search != "" {
		search = strings.ReplaceAll(search, " ", "")
		url = "https://www.googleapis.com/books/v1/volumes?q=" + search + "&printType=books&langRestrict=pt"
	} else {
		url = "https://www.googleapis.com/books/v1/volumes?q=''&printType=books&langRestrict=pt"
	}

	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if response.StatusCode != http.StatusOK {
		c.JSON(response.StatusCode, gin.H{"error": fmt.Sprintf("error: %v", response.StatusCode)})
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var responseGoogle models.SearchBooksModel
	err = json.Unmarshal(body, &responseGoogle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responseGoogle)
}
