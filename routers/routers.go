package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ianbruce/todo/handlers"
	"github.com/ianbruce/todo/injects"

)

type RouteData struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

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

func createRoutes(appCtx injects.AppContainer) []RouteData {
	return []RouteData{
		RouteData{
			"Index",
			"GET",
			"/",
			handlers.Index(&appCtx),
		},

		RouteData{
			"AddList",
			"POST",
			"/lists",
			handlers.AddList(&appCtx),
		},

		RouteData{
			"AddTask",
			"POST",
			"/list/{id}/tasks",
			handlers.AddTask(&appCtx),
		},

		RouteData{
			"GetList",
			"GET",
			"/list/{id}",
			handlers.GetList(&appCtx),
		},

		RouteData{
			"PutTask",
			"POST",
			"/list/{id}/task/{taskId}/complete",
			handlers.UpdateTaskCompletion(&appCtx),
		},

		RouteData{
			"SearchLists",
			"GET",
			"/lists",
			handlers.SearchLists(&appCtx),
		},
	}
}
