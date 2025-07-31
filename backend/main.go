package main

import (
	"to-do-app/database"
	"to-do-app/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.Connect()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("api/todos", handler.GetTodos)
	e.POST("/api/todos", handler.AddTodo)
	e.DELETE("/api/todos/:id", handler.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
