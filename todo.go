package main

import (
	_ "encoding/json"
)

func Create(message string) Todo {
	var id = CreateTodoEntity(message)
	return GetTodo(id)
}
func GetAll() []Todo {
	return GetAllEntities()
}
func Update(id int, todo Todo) Todo {
	original := GetTodo(id)
	if todo.Status == "" {
		todo.Status = original.Status
	}
	if todo.Message == "" {
		todo.Message = original.Message
	}
	UpdateEntity(id, todo)
	return GetTodo(id)
}
func GetTodo(Id int) Todo {
	return GetTodoEntity(Id)
}
