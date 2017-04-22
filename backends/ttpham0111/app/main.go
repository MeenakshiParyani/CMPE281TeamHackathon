package main

import (
	"ttpham0111/app/web"
	"ttpham0111/app/web/middleware"
)

func main() {
	datastore := middleware.LocalDatastore{}
	web.StartServer(&datastore)
}
