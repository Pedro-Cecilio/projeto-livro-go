package main

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/database"
	"github.com/Pedro-Cecilio/projeto-livro-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.CreateConnection()
	r := gin.Default()
	routes.UserRoutes(r)
	routes.BookRoutes(r)
	r.Run(":3333")

	database.CreateConnection()
}
