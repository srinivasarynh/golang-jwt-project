package main

import (
	"fmt"
	"golang-jwt-project/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.Default()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "access to api-1",
		})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "access to api-2",
		})
	})

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("server error", err)
	}
	fmt.Println("server running at port 8000")
}
