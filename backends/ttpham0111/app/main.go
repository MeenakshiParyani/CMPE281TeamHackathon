package main

import (
	"ttpham0111/app/web"
	"ttpham0111/app/web/middleware"
)

func main() {
	datastore := middleware.NewMgoDatastore("mongodb://mongodb/cmpe281")
	defer datastore.Close()

	numBaristas := 8
	web.StartServer(datastore, numBaristas)
}
