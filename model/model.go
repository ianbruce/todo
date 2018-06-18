package model

type Task struct {
	ID           string `dynamodbav:"taskID"`
	Name         string `dynamodbav:"taskName"`
	Completed    bool   `dynamodbav:"taskCompleted"`
}

type TodoList struct {
	ID           string `dynamodbav:"listID"`
	Name         string `dynamodbav:"listName"`
	Description  string `dynamodbav:"listDescription"`
	Tasks        []Task `dynamodbav:"tasks"`
}
