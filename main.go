package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"os"
	routes "user-athentication-golang"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.new()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	//second Api

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted For Api-1"})
	})

	// API-1
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)

}
