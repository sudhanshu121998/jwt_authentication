package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sudhanshu121998/authentication_module/pkg/controller"
)

func AuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", controller.LoginUser)
		authGroup.POST("/register", controller.RegisterUser)
	}
}
