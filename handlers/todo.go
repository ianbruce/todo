package handlers

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ianbruce/todo/model"
	"github.com/ianbruce/todo/injects"
)

func Index(appCtn *injects.AppContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write([]byte("howdy!"))
	}
}

func AddList(appCtn *injects.AppContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var myList model.TodoList

		b, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal([]byte(b), &myList)

		if err == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
	}
}

func AddTask(appCtn *injects.AppContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
	}
}

func GetList(appCtn *injects.AppContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		// validate the incoming request
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		list, getErr := appCtn.DB.GetList(id)

		// if there was error on getting the list, it didn't exist
		if getErr != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		marshalled, _ := json.Marshal(list)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(marshalled)
	}
}

func PutTask(appCtn *injects.AppContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}

func SearchLists(appCtn *injects.AppContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		searchStringField := r.Form["searchString"]
		limitField := r.Form["limit"]

		var searchString string
		if len(searchStringField) == 0 {
			searchString = ""
		} else {
			searchString = searchStringField[0]
		}

		var limit int
		if len(limitField) == 0 {
			limit = 50
		} else {
			parsedLimit, parseErr := strconv.ParseInt(limitField[0], 10, 32)
			if parseErr != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("limit parameter is not an integer"))
				return
			}
			limit = int(parsedLimit)
		}

		lists, getErr := appCtn.DB.GetLists(searchString, limit)

		if getErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(getErr.Error()))
			return
		}

		marshalled, marshalErr := json.Marshal(lists)

		if marshalErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(marshalled)
	}
}
