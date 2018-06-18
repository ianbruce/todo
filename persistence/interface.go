package persistence

import (
  "github.com/ianbruce/todo/model"
)

type TodoDatabase interface {
  GetList(id string) (model.TodoList, error)
  GetLists(query string, limit int) ([]model.TodoList, error)
  CreateList(list model.TodoList) error
  CreateTask(listId string, task model.Task) error
  UpdateTaskStatus(listId string, taskId string) error
}
