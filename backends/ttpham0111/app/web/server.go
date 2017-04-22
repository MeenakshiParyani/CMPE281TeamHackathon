package web

import (
	"log"
	"net/http"
	"ttpham0111/app/web/middleware"
)

func StartServer(datastore middleware.Datastore) {
	router := NewRouter(datastore)
	log.Fatal(http.ListenAndServe(":5000", router))
}
