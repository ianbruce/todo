package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ianbruce/todo/middleware"
	"github.com/ianbruce/todo/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router	 := mux.NewRouter().StrictSlash(true)
	for _, route := range myRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middleware.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var myRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},

	Route{
		"AddList",
		"POST",
		"/lists",
		handlers.AddList,
	},

	Route{
		"AddTask",
		"POST",
		"/list/{id}/tasks",
		handlers.AddTask,
	},

	Route{
		"GetList",
		"GET",
		"/list/{id}",
		handlers.GetList,
	},

	Route{
		"PutTask",
		"POST",
		"/list/{id}/task/{taskId}/complete",
		handlers.PutTask,
	},

	Route{
		"SearchLists",
		"GET",
		"/lists",
		handlers.SearchLists,
	},
}
