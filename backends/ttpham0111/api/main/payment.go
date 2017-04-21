package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order-id"]

	err, order := datastore.GetOrder(orderID)

	if err != nil {
		switch e := err.(type) {
		case NoResultFound:
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(APIError{e.Error()})
			return
		default:
			panic(e)
		}
	}

	if order.status != PLACED {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIError{"Payment Rejected"})
		return
	}

	order.SetOrderStatus(PAID)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		panic(err)
	}
}
