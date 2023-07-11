package routes

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/controllers"
	"github.com/Pedro-Cecilio/projeto-livro-go/middleware"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine){
	book := router.Group("/book")
	book.Use(middleware.CheckIfAuthorizationExist(), middleware.CheckIfAuthorizationIsValid())
	// book.Use(middleware.CheckIfAuthorizationIsValid())

	book.GET("/", controllers.SearchBooksController)
	book.GET("/:id", controllers.GetBookController)
	book.POST("/BookRead", controllers.InsertBookReadController)

}