package persistence

import (
  "github.com/ianbruce/todo/model"
)

type TodoDatabase interface {
  GetList(id string) model.TodoList
  GetLists(query string, limit int) []model.TodoList
  CreateList(list model.TodoList)
  CreateTask(listId string, task model.Task)
  UpdateTaskStatus(listId string, taskId string)
}
