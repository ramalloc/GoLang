package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ramalloc/go-jwt/routes"
)

func main()  {
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}
	router := gin.New()

	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success":"Access Granted For api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success":"Access Granted For api-2"})
	})

	router.Run(":" + port)
}