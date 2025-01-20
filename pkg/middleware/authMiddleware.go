package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudhanshu121998/authentication_module/pkg/helpers"
)

func Authenticate(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")

	if clientToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No AuthHeader provided"})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(clientToken)

	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	c.Set("email", claims.Email)
	c.Set("first_name", claims.FirstName)
	c.Set("last_name", claims.LastName)
	c.Set("uid", claims.Uid)
	c.Set("user_type", claims.UserType)
	c.Next()
	
}
