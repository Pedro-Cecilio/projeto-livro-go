package routes

import (
	"github.com/Pedro-Cecilio/projeto-livro-go/controllers"
	"github.com/Pedro-Cecilio/projeto-livro-go/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/", controllers.CreateUserController)
	router.POST("login", controllers.UserLoginController)
	router.POST("/logout", middleware.CheckIfAuthorizationExist(), middleware.CheckIfAuthorizationIsValid(), controllers.UserLogoutController)
	
}
