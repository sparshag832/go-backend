package main

import (
	"gincrud/config"
	"gincrud/routes"

	_ "gincrud/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	config.ConnectMongoDB()

	// Create a new Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
