package web

import (
	"fmt"
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
		go middleware.NewBarista(rand.Float32()+0.5, queue, datastore).Start()
	}

	router := NewRouter(datastore, queue)
	fmt.Println("Starting server on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
