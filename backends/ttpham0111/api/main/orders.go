// TODO: Replace panics with properly handled responses
// TODO: Manage order state
// TODO: Prevent init/replacing of fields used by server

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	err, orders := datastore.GetOrders()
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(&orders); err != nil {
		panic(err)
	}
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
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

	if err := json.NewEncoder(w).Encode(&order); err != nil {
		panic(err)
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err, order := datastore.CreateOrder(&newOrder)
	if err != nil {
		panic(err)
	}

	order.SetOrderStatus(PLACED)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&order); err != nil {
		panic(err)
	}
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order-id"]

	var newOrder Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err, order := datastore.UpdateOrder(orderID, &newOrder)
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(&order); err != nil {
		panic(err)
	}
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order-id"]
	err := datastore.DeleteOrder(orderID)
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

	w.WriteHeader(http.StatusNoContent)
}
