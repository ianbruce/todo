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

var attributeNameMap = map[string]*string{
  "#owner": aws.String("owner"),
  // "#pub": aws.String("public"),

}

func (service *DynamoTodo) GetList(id string) (model.TodoList, error) {
  var returnList model.TodoList

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
    return returnList, getErr
  }

  unmarshalErr := dynamodbattribute.UnmarshalMap(resp.Item, &returnList)

  return returnList, unmarshalErr
}

func (service *DynamoTodo) GetLists(query string, limit int) ([]model.TodoList, error) {
  var returnList []model.TodoList

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
      TableName: aws.String("todos"),
    }
  }

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

  updateExpression := "SET tasks = list_append(tasks, :newTask)"
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
    },
    ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
      ":newTask": {
        L: []*dynamodb.AttributeValue{
          {
            M: newTaskAV,
          },
        },
      },
    },
    UpdateExpression: &updateExpression,
    },
  )

  return updateErr
}

func (service *DynamoTodo) UpdateTaskStatus(listId string, taskId string) error {
  updateVal, updateErr := service.DB.UpdateItem(&dynamodb.UpdateItemInput{
    ExpressionAttributeNames: attributeNameMap,
    Key: map[string]*dynamodb.AttributeValue{
      "#owner": {
        S: aws.String("public"),
      },
      "listID": {
        S: aws.String(listId),
      },
      "taskID": {
        S: aws.String(listId),
      },
    },
    // ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
    //   ":newStatus": {
    //     BOOL: aws.Bool(true),
    //   },
    // },
    UpdateExpression: aws.String("SET completed = true"),
  },)

  fmt.Println(updateVal.Attributes)

  return updateErr
}
