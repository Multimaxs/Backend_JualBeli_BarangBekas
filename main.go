package main

import (
	"backend-jual-beli/config"
	"backend-jual-beli/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	config.SetupDatabase()

	// Create a Gin router
	router := gin.Default()

	// Set up routing
	routes.SetupRoutes(router)

	// Run the application on port 8080
	router.Run(":8080")

	// Initialize JWT
	config.InitConfig()

}
