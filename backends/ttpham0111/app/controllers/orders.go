// TODO: Replace panics with properly handled responses
// TODO: Keep it DRY with error handling middleware

package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
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
	err, order := getOrder(orderID, controller.datastore, w, r)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(&order); err != nil {
		panic(err)
	}
}

func (controller *OrdersController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	newOrder := models.Order{
		ID:       bson.NewObjectId(),
		Location: orderRequest.Location,
		Items:    orderRequest.Items,
	}
	newOrder.SetOrderStatus(models.OrderPlaced, r.Host)

	err, order := controller.datastore.CreateOrder(&newOrder)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		panic(err)
	}
}

func (controller *OrdersController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order-id"]
	err, order := getOrder(orderID, controller.datastore, w, r)
	if err != nil {
		return
	} else if order.Status != models.OrderPlaced {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIError{"Order Update Rejected."})
		return
	}

	var orderRequest models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	order.Location = orderRequest.Location
	order.Items = orderRequest.Items

	err, order = controller.datastore.UpdateOrder(orderID, order)
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(&order); err != nil {
		panic(err)
	}
}

func (controller *OrdersController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order-id"]
	if err, order := getOrder(orderID, controller.datastore, w, r); err != nil {
		return
	} else if order.Status != models.OrderPlaced {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIError{"Order Cancelling Rejected."})
		return
	}

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

func getOrder(orderID string, datastore middleware.Datastore, w http.ResponseWriter, r *http.Request) (error, *models.Order) {
	err, order := datastore.GetOrder(orderID)
	if err != nil {
		switch e := err.(type) {
		case middleware.NoResultFound:
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(APIError{e.Error()})
			return e, nil
		default:
			panic(e)
		}
	}

	return nil, order
}
