// TODO: Replace panics with properly handled responses
// TODO: Manage order state
// TODO: Prevent init/replacing of fields used by server

package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"ttpham0111/app/models"
	"ttpham0111/app/web/middleware"
)

type OrdersController struct {
	datastore middleware.Datastore
}

func NewOrdersController(datastore middleware.Datastore) *OrdersController {
	return &OrdersController{
		datastore,
	}
}

func (controller *OrdersController) GetOrders(w http.ResponseWriter, r *http.Request) {
	err, orders := controller.datastore.GetOrders()
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(&orders); err != nil {
		panic(err)
	}
}

func (controller *OrdersController) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order-id"]

	err, order := controller.datastore.GetOrder(orderID)
	if err != nil {
		switch e := err.(type) {
		case middleware.NoResultFound:
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

func (controller *OrdersController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err, order := controller.datastore.CreateOrder(&newOrder)
	if err != nil {
		panic(err)
	}

	order.SetOrderStatus(models.OrderPlaced, r.Host)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&order); err != nil {
		panic(err)
	}
}

func (controller *OrdersController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order-id"]

	var newOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err, order := controller.datastore.UpdateOrder(orderID, &newOrder)
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(&order); err != nil {
		panic(err)
	}
}

func (controller *OrdersController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order-id"]
	err := controller.datastore.DeleteOrder(orderID)
	if err != nil {
		switch e := err.(type) {
		case middleware.NoResultFound:
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(APIError{e.Error()})
			return
		default:
			panic(e)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
