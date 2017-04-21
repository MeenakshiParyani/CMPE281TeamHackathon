package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		handler := JsonifyResponse(route.HandlerFunc)

		router.
			Methods(route.Methods...).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
