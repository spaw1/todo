package main

import (
	_ "encoding/json"
)

func Create(message string) int64 {
	return CreateTodoEntity(message)
}
func GetAll() []Todo {
	return GetAllEntities()
}

func GetTodo(Id string) Todo {
	return GetTodoEntity(Id)
}
func MarkAsComplete(Id string) bool {
	return MarkAsCompleteEntity(Id)
}

func Delete(Id string) bool {
	return DeleteEntity(Id)
}
