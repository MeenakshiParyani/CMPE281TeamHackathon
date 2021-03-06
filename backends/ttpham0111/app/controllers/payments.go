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
	err, order := getOrder(orderID, controller.datastore, w, r)
	if err != nil {
		return
	}

	if order.Status != models.OrderPlaced {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIError{"Payment Rejected"})
		return
	}

	order.SetOrderStatus(models.OrderPaid, "")
	err, order = controller.datastore.UpdateOrder(order.ID.Hex(), order)
	if err != nil {
		panic(err)
	}

	controller.queue <- order

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		panic(err)
	}
}
