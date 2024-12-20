package routes

import (
	"backend-jual-beli/controllers"
	"backend-jual-beli/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Grouping API routes under "/api"
	api := router.Group("/api")

	// Public routes
	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)

	// Protected routes (require authentication)
	api.Use(middlewares.AuthMiddleware())

	// User routes
	api.GET("/users", controllers.GetUsers)             // Get all users
	api.GET("/users/:id", controllers.GetUserByID)      // Get user by ID
	api.PUT("/users/:id", controllers.UpdateUser)       // Update user by ID
	api.DELETE("/users/:id", controllers.DeleteUser)    // Delete user by ID

	// Item routes
	api.POST("/items", controllers.CreateItem)          // Create a new item
	api.GET("/items", controllers.GetItems)             // Get all items
	api.GET("/items/:id", controllers.GetItemByID)      // Get item by ID
	api.PUT("/items/:id", controllers.UpdateItem)       // Update item by ID
	api.DELETE("/items/:id", controllers.DeleteItem)    // Delete item by ID

	// Order routes
	api.POST("/orders", controllers.CreateOrder)        // Create a new order
	api.GET("/orders", controllers.GetOrders)           // Get all orders
	api.GET("/orders/:id", controllers.GetOrderByID)    // Get order by ID
	api.PUT("/orders/:id", controllers.UpdateOrder)     // Update order by ID
	api.DELETE("/orders/:id", controllers.DeleteOrder)  // Delete order by ID
}
