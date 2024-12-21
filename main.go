package main

import (
	"todo-app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Register routes
	routes.RegisterRoutes(e)

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
