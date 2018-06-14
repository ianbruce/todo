package model

import (
  "database/sql"
)

type Model struct {
  store *sql.DB
}

type Task struct {
	id           string `db: "id"`
	name         string `db: "name"`
	completed    bool   `db: "completed"`
}

type TodoList struct {
	id           string `db: "id"`
	name         string `db: "name"`
	description  string `db: "description"`
	tasks        []Task `db: "tasks"`
}
