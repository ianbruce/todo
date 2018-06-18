package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ianbruce/todo/handlers"
	"github.com/ianbruce/todo/injects"

)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


func NewRouter(appCtx injects.AppContainer) *mux.Router {
	router	 := mux.NewRouter().StrictSlash(true)
	for _, route := range createRoutes(appCtx) {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func createRoutes(appCtx injects.AppContainer) Routes {
	return Routes{
		Route{
			"Index",
			"GET",
			"/",
			handlers.Index(&appCtx),
		},

		Route{
			"AddList",
			"POST",
			"/lists",
			handlers.AddList(&appCtx),
		},

		Route{
			"AddTask",
			"POST",
			"/list/{id}/tasks",
			handlers.AddTask(&appCtx),
		},

		Route{
			"GetList",
			"GET",
			"/list/{id}",
			handlers.GetList(&appCtx),
		},

		Route{
			"PutTask",
			"POST",
			"/list/{id}/task/{taskId}/complete",
			handlers.PutTask(&appCtx),
		},

		Route{
			"SearchLists",
			"GET",
			"/lists",
			handlers.SearchLists(&appCtx),
		},
	}
}
