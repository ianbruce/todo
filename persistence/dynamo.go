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

func (service *DynamoTodo) GetList(id string) model.TodoList {
  var returnList model.TodoList

  dbInput := &dynamodb.GetItemInput{
    Key: map[string]*dynamodb.AttributeValue{
        "listID": {
            S: aws.String(id),
        },
    },
    TableName: aws.String(service.TableName),
  }

  resp, _ := service.DB.GetItem(dbInput)

  dynamodbattribute.UnmarshalMap(resp.Item, &returnList)

  return returnList
}

func (service *DynamoTodo) GetLists(query string, limit int) []model.TodoList {
  var returnList []model.TodoList

  service.DB.ScanPages(&dynamodb.ScanInput{
      TableName: aws.String(service.TableName),
    }, func(page *dynamodb.ScanOutput, last bool) bool {
        lists := []model.TodoList{}

        err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &lists)

        if err != nil {
             panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
        }

        returnList = append(returnList, lists...)

        return true // keep paging
      },
  )

  return returnList
}

func (service *DynamoTodo) CreateList(list model.TodoList) {
  av, err := dynamodbattribute.MarshalMap(list)

  if err != nil {
      panic(fmt.Sprintf("failed to DynamoDB marshal list, %v", err))
  }

  _, err = service.DB.PutItem(&dynamodb.PutItemInput{
      TableName: aws.String(service.TableName),
      Item:      av,
  })

  if err != nil {
      panic(fmt.Sprintf("failed to put list to DynamoDB, %v", err))
  }
}

func (service *DynamoTodo) CreateTask(listId string, task model.Task) {

}

func (service *DynamoTodo) UpdateTaskStatus(listId string, taskId string) {
  _, err := service.DB.UpdateItem(&dynamodb.UpdateItemInput{
      Key: map[string]*dynamodb.AttributeValue{
          "listID": {
              S: aws.String(listId),
          },
          "taskID": {
              S: aws.String(taskId),
          },
      },
      TableName: aws.String(service.TableName),
    },
  )

  if err != nil {
    panic(fmt.Sprintf("couldn't update task %s completion status!", taskId))
  }
}
