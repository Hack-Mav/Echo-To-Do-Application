package routes

import (
	"todo-app/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/tasks", handlers.GetTasks)
	e.GET("/tasks/:id", handlers.GetTask)
	e.POST("/tasks", handlers.CreateTask)
	e.PUT("/tasks/:id", handlers.UpdateTask)
	e.DELETE("/tasks/:id", handlers.DeleteTask)
}
