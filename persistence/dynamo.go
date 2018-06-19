package persistence

import (
  "fmt"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

  "github.com/ianbruce/todo/model"
)

type DynamoTodo struct {
  DB        *dynamodb.DynamoDB
  TableName string
}

type TodoMap struct {
  ID           string `dynamodbav:"listID"`
	Name         string `dynamodbav:"listName"`
	Description  string `dynamodbav:"listDescription"`
	Tasks        map[string]*model.Task `dynamodbav:"tasks"`
}

func (tm *TodoMap) ToTodoList() model.TodoList {
  var taskList []model.Task
  for _, v := range tm.Tasks {
    taskList = append(taskList, *v)
  }
  return model.TodoList{
    ID: tm.ID,
    Name: tm.Name,
    Description: tm.Description,
    Tasks: taskList,
  }
}

func todoListToTodoMap(tl model.TodoList) TodoMap {
  taskMap := make(map[string]*model.Task)
  for _, v := range tl.Tasks {
    taskMap[v.ID] = &v
  }
  return TodoMap{
    ID: tl.ID,
    Name: tl.Name,
    Description: tl.Description,
    Tasks: taskMap,
  }
}

var attributeNameMap = map[string]*string{
  "#owner": aws.String("owner"),
}

func (service *DynamoTodo) GetList(id string) (model.TodoList, error) {
  var returnedMap TodoMap

  dbInput := &dynamodb.GetItemInput{
    Key: map[string]*dynamodb.AttributeValue{
      "owner": {
        S: aws.String("public"),
      },
      "listID": {
        S: aws.String(id),
      },
    },
    TableName: aws.String(service.TableName),
  }

  resp, getErr := service.DB.GetItem(dbInput)

  if getErr != nil {
    return returnedMap.ToTodoList(), getErr
  }

  unmarshalErr := dynamodbattribute.UnmarshalMap(resp.Item, &returnedMap)

  return returnedMap.ToTodoList(), unmarshalErr
}

func (service *DynamoTodo) GetLists(query string, limit int) ([]model.TodoList, error) {
  var returnedMapList []TodoMap

  var queryInput *dynamodb.QueryInput
  if query == "" {
    queryInput = &dynamodb.QueryInput{
      ExpressionAttributeNames: attributeNameMap,
      ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
        ":public": {
          S: aws.String("public"),
        },
      },
      KeyConditionExpression: aws.String("#owner = :public"),
      Limit: aws.Int64(int64(limit)),
      TableName: aws.String("todos"),
    }
  } else {
    queryInput = &dynamodb.QueryInput{
      ExpressionAttributeNames: attributeNameMap,
      ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
        ":prefix": {
          S: aws.String(query),
        },
        ":public": {
          S: aws.String("public"),
        },
      },
      KeyConditionExpression: aws.String("#owner = :public AND begins_with(listID, :prefix)"),
      Limit: aws.Int64(int64(limit)),
      TableName: aws.String("todos"),
    }
  }

  fmt.Println(queryInput)

  lists, queryErr := service.DB.Query(queryInput)

  if queryErr != nil {
    var returnList []model.TodoList
    return returnList, queryErr
  }

  unmarshalErr := dynamodbattribute.UnmarshalListOfMaps(lists.Items, &returnedMapList)

  var returnList []model.TodoList
  for _, v := range returnedMapList {
      returnList = append(returnList, v.ToTodoList())
  }

  return returnList, unmarshalErr
}

func (service *DynamoTodo) CreateList(list model.TodoList) error {
  todoMap := todoListToTodoMap(list)
  av, marhsalErr := dynamodbattribute.MarshalMap(todoMap)

  if marhsalErr != nil {
    return marhsalErr
  }

  av["owner"] = &dynamodb.AttributeValue{
    S: aws.String("public"),
  }

  fmt.Println(av)

  _, putErr := service.DB.PutItem(&dynamodb.PutItemInput{
    TableName: aws.String(service.TableName),
    Item:      av,
  })

  return putErr
}

func (service *DynamoTodo) CreateTask(listId string, task model.Task) error {
  newTaskAV, marshalErr := dynamodbattribute.MarshalMap(task)

  if marshalErr != nil {
    return marshalErr
  }

  _, updateErr := service.DB.UpdateItem(&dynamodb.UpdateItemInput{
    TableName: aws.String(service.TableName),
    ExpressionAttributeNames: attributeNameMap,
    Key: map[string]*dynamodb.AttributeValue{
      "#owner": {
        S: aws.String("public"),
      },
      "listID": {
          S: aws.String(listId),
      },
      "tasks": {
        S: aws.String("tasks"),
      },
    },
    ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
      ":newTask": {
        L: []*dynamodb.AttributeValue{
          {
            M: newTaskAV,
          },
        },
      },
      ":taskID": {
        S: aws.String(task.ID),
      },
    },
    UpdateExpression: aws.String("SET :taskID = :newTask"),
    },
  )

  return updateErr
}

func (service *DynamoTodo) UpdateTaskStatus(listId string, taskId string, completion bool) error {
  _, updateErr := service.DB.UpdateItem(&dynamodb.UpdateItemInput{
    TableName: aws.String(service.TableName),
    Key: map[string]*dynamodb.AttributeValue{
      "owner": {
        S: aws.String("public"),
      },
      "listID": {
        S: aws.String(listId),
      },
      // "tasks": {
      //   S: aws.String("tasks"),
      // },
      // taskId: {
      //   S: aws.String(taskId),
      // },
    },
    ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
      ":newStatus": {
        BOOL: aws.Bool(completion),
      },
    },
    UpdateExpression: aws.String(fmt.Sprintf("SET tasks.%s.taskCompleted = :newStatus", taskId)),
  },)

  return updateErr
}
