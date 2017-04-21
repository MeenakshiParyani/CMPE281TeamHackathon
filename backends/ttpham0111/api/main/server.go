package main

import (
	"log"
	"net/http"
)

var datastore = LocalDatastore{}

type APIError struct {
	Error string `json:"error"`
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", NewRouter()))
}
