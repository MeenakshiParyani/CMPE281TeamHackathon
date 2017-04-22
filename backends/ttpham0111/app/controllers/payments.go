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
}

func NewPaymentController(datastore middleware.Datastore) *PaymentController {
	return &PaymentController{
		datastore,
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
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		panic(err)
	}
}
