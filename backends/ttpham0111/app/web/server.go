package web

import (
	"log"
	"math/rand"
	"net/http"
	"ttpham0111/app/models"
	"ttpham0111/app/web/middleware"
)

func StartServer(datastore middleware.Datastore, numBaristas int) {

	// Start Baristas
	queue := make(chan *models.Order, numBaristas)
	for i := 0; i < numBaristas; i++ {
		go models.Barista{queue, rand.Float32() + 0.5}.Start()
	}

	router := NewRouter(datastore, queue)
	log.Fatal(http.ListenAndServe(":5000", router))
}
