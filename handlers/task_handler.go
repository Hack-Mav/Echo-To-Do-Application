package handlers

import (
	"net/http"
	"strconv"
	"todo-app/models"

	"github.com/labstack/echo/v4"
)

// Simulated database
var tasks = []models.Task{}
var nextID = 1

// GetTasks retrieves all tasks
func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

// GetTask retrieves a task by ID
func GetTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, t := range tasks {
		if t.ID == id {
			return c.JSON(http.StatusOK, t)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

// CreateTask creates a new task
func CreateTask(c echo.Context) error {
	t := new(models.Task)
	if err := c.Bind(t); err != nil {
		return err
	}
	t.ID = nextID
	nextID++
	tasks = append(tasks, *t)
	return c.JSON(http.StatusCreated, t)
}

// UpdateTask updates an existing task
func UpdateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range tasks {
		if t.ID == id {
			if err := c.Bind(&tasks[i]); err != nil {
				return err
			}
			tasks[i].ID = id // Ensure the ID is not updated
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

// DeleteTask deletes a task by ID
func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "task deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}
