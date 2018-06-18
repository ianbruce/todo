package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/patrickmn/go-cache"

	"github.com/ianbruce/todo/routers"
	"github.com/ianbruce/todo/persistence"
	"github.com/ianbruce/todo/injects"

)

func main() {
	myDynamo := dynamodb.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

	dynamoTodo := persistence.DynamoTodo{
		DB: myDynamo,
		TableName: "todos",
	}

	cachingProxy := persistence.CacheProxy{
		Proxied: &dynamoTodo,
		Cache: cache.New(5*time.Minute, 10*time.Minute),
	}

	appContainer := injects.AppContainer{
		DB: &cachingProxy,
	}

	router := routers.NewRouter(appContainer)

	log.Fatal(http.ListenAndServe(":8080", router))
}
