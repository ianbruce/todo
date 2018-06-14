package main

import (
	"log"
	"net/http"
	"github.com/ianbruce/todo/routers"
)

func main() {
	log.Printf("Server started")

	router := routers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
