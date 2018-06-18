package injects

import (
  "github.com/ianbruce/todo/persistence"
)

type AppContainer struct {
  DB persistence.TodoDatabase
}
