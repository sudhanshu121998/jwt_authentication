package main

import (
	"os"

	"github.com/gin-gonic/gin"

	routes "github.com/sudhanshu121998/authentication_module/pkg/router"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)

}
