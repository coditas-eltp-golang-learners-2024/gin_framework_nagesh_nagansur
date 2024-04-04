// main.go
package main

import (
	"Project/middleware"
	"Project/routes"
)

func main() {
	// Initialize the Gin router
	router := routes.SetupRouter()
	router.Use(middleware.LoggerMiddleware())
	// Start the server
	router.Run(":8080")
}
