package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"ttpham0111/app/models"
	"ttpham0111/app/web/middleware"
)

type PaymentController struct {
	datastore middleware.Datastore
	queue     chan *models.Order
}

func NewPaymentController(datastore middleware.Datastore, queue chan *models.Order) *PaymentController {
	return &PaymentController{
		datastore,
		queue,
	}
}

func (controller *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
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

	if order.Status != models.OrderPlaced {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIError{"Payment Rejected"})
		return
	}

	order.SetOrderStatus(models.OrderPaid)
	controller.queue <- order

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		panic(err)
	}
}
