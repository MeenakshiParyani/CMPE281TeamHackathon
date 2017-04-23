package main

import (
	"os"
	"ttpham0111/app/web"
	"ttpham0111/app/web/middleware"
)

func main() {
	var datastore middleware.Datastore
	if dbUrl := os.Getenv("DB_URL"); dbUrl != "" {
		datastore = middleware.NewMgoDatastore(dbUrl)
	} else {
		datastore = &middleware.LocalDatastore{}
	}
	defer datastore.Close()

	numBaristas := 8
	web.StartServer(datastore, numBaristas)
}
