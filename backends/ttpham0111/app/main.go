package main

import (
	"ttpham0111/app/web"
	"ttpham0111/app/web/middleware"
)

func main() {
	datastore := middleware.LocalDatastore{}
	numBaristas := 8
	web.StartServer(&datastore, numBaristas)
}
