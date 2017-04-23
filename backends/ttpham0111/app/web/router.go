package web

import (
	"github.com/gorilla/mux"
	"ttpham0111/app/models"
	"ttpham0111/app/web/middleware"
)

func NewRouter(datastore middleware.Datastore, queue chan *models.Order) *mux.Router {
	router := mux.NewRouter()

	routes := GetRoutes(datastore, queue)
	for _, route := range routes {
		handler := middleware.JsonifyResponse(route.HandlerFunc)

		router.
			Methods(route.Methods...).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
