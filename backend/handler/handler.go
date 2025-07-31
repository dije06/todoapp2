package handler

import (
	"net/http"
	"strconv"
	"to-do-app/model"
	"to-do-app/repository"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	todos, err := repository.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch todos"})
	}
	return c.JSON(http.StatusOK, todos)
}

func AddTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdTodo, err := repository.AddToDo(todo.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, createdTodo)
}

func DeleteTodo(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	err = repository.DeleteToDo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete todo"})
	}
	return c.NoContent(http.StatusNoContent)
}
