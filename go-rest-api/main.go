package main

import (
	"go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	//database.ConnectDatabase()

	// Initialize the Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Run the server
	router.Run(":2313")
}
