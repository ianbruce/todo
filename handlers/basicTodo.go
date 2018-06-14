package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"github.com/ianbruce/todo/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func AddList(w http.ResponseWriter, r *http.Request) {
	var myList model.TodoList
	b, readErr := ioutil.ReadAll(r.Body)
	if readErr == nil {
		err := json.Unmarshal([]byte(b), &myList)
		if err == nil {
			fmt.Println("here's the list u added:")
			fmt.Println("%+v", myList)
		}
	} else {
		fmt.Println("couldn't read in body of request")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PutTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func SearchLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
