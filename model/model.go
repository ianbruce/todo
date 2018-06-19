package model

type Task struct {
	ID           string `dynamodbav:"taskID", json: "id"`
	Name         string `dynamodbav:"taskName", json: "name"`
	Completed    bool   `dynamodbav:"taskCompleted", json: "completed"`
}

type TodoList struct {
	ID           string `dynamodbav:"listID", json: "id"`
	Name         string `dynamodbav:"listName", json: "name"`
	Description  string `dynamodbav:"listDescription", json: "description"`
	Tasks        []Task `dynamodbav:"tasks", json: "tasks"`
}

type CompletedTask struct {
  Completed    bool   `dynamodbav:"completed"`
}
