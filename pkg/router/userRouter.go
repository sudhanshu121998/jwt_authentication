package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sudhanshu121998/authentication_module/pkg/controller"
	"github.com/sudhanshu121998/authentication_module/pkg/middleware"
)

func UserRoutes(r *gin.Engine) {
	r.Use(middleware.Authenticate)
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", controller.GetUsers)
		userGroup.GET("/:id", controller.GetUser)
	}
}
