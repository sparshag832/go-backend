package routes

import (
	"gincrud/controllers"

	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Swagger route
	router.GET("/swagger/*filepath", func(c *gin.Context) {
		httpSwagger.WrapHandler.ServeHTTP(c.Writer, c.Request)
	})

	// Define API routes
	router.GET("/users", func(c *gin.Context) {
		controllers.GetUsers(c)
	})

	router.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(c)
	})
}
