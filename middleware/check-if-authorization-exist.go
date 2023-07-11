package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckIfAuthorizationExist() gin.HandlerFunc{
	return func (c *gin.Context){
		if authorization := c.GetHeader("Authorization"); authorization == ""{
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
		
		
		
	}
}