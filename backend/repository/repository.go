package repository

import (
	"to-do-app/database"
	"to-do-app/model"
)

func GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	result := database.DB.Find(&todos)
	return todos, result.Error
}

func AddToDo(title string) (model.Todo, error) {
	newTodo := model.Todo{
		Title: "testing_" + title, // keep your logic or remove "testing_"
	}
	result := database.DB.Create(&newTodo)
	return newTodo, result.Error
}

func DeleteToDo(id int) error {
	result := database.DB.Delete(&model.Todo{}, id)
	return result.Error
}
