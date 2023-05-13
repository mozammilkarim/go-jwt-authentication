package main

import (
	"os"

	"github.com/gin-gonic/gin"
	routes "github.com/syed/go-jwt-authentication/routes"
)

func main() {
	// here without loading,read port value from .env
	//  because os is directly at same level
	port := os.Getenv("PORT")

	if port != "" {
		port = "8000"
	}
	router := gin.New()      // create route
	router.Use(gin.Logger()) // to log routes accessed
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	//what is its use
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"}) // setting up headers
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-2"})
	})
	// server started
	router.Run(":" + port)
}
