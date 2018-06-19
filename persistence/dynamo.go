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

func (service *DynamoTodo) GetList(id string) (model.TodoList, error) {
  var returnList model.TodoList

  dbInput := &dynamodb.GetItemInput{
    Key: map[string]*dynamodb.AttributeValue{
        "listID": {
            S: aws.String(id),
        },
    },
    TableName: aws.String(service.TableName),
  }

  resp, getErr := service.DB.GetItem(dbInput)

  if getErr != nil {
    return returnList, getErr
  }

  unmarshalErr := dynamodbattribute.UnmarshalMap(resp.Item, &returnList)

  return returnList, unmarshalErr
}

func (service *DynamoTodo) GetLists(query string, limit int) ([]model.TodoList, error) {
  var returnList []model.TodoList
  queryInput := &dynamodb.QueryInput{
    TableName: aws.String(service.TableName),
  }

  queryInput = queryInput.SetExpressionAttributeValues(map[string]*dynamodb.AttributeValue{
    ":prefix": {
        S: aws.String(query),
    },
  })

  queryInput = queryInput.SetKeyConditionExpression("begins_with(listID, :prefix)")

  fmt.Println(queryInput)

  lists, queryErr := service.DB.Query(queryInput)

  if queryErr != nil {
    return returnList, queryErr
  }

  unmarshalErr := dynamodbattribute.UnmarshalListOfMaps(lists.Items, &returnList)

  return returnList, unmarshalErr
}

func (service *DynamoTodo) CreateList(list model.TodoList) error {
  av, marhsalErr := dynamodbattribute.MarshalMap(list)

  if marhsalErr != nil {
    return marhsalErr
  }

  _, putErr := service.DB.PutItem(&dynamodb.PutItemInput{
      TableName: aws.String(service.TableName),
      Item:      av,
  })

  return putErr
}

// TODO: implement CreateTask
func (service *DynamoTodo) CreateTask(listId string, task model.Task) error {
  newTaskAV, marshalErr := dynamodbattribute.MarshalMap(task)

  _, updateErr := service.DB.UpdateItem(&dynamodb.UpdateItemInput{
      TableName: aws.String(service.TableName),
      Key: map[string]*dynamodb.AttributeValue{
          "listID": {
              S: aws.String(listId),
          },
      },
      ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
        ":newVal": {
          BOOL: aws.Bool(true),
        },
      },
      UpdateExpression: "ADD completed = :newStatus",
    },
  )

  return updateErr
}

func (service *DynamoTodo) UpdateTaskStatus(listId string, taskId string) error {
  _, updateErr := service.DB.UpdateItem(&dynamodb.UpdateItemInput{
      Key: map[string]*dynamodb.AttributeValue{
          "listID": {
              S: aws.String(listId),
          },
          "taskID": {
              S: aws.String(listId),
          },
      },
      ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
        ":newStatus": {
          BOOL: aws.Bool(true),
        },
      },
      UpdateExpression: "SET completed = :newStatus",
    },
  )

  return updateErr
}
