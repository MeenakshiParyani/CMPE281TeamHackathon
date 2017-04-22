package web

import (
	"encoding/json"
	"net/http"
	"ttpham0111/app/controllers"
	"ttpham0111/app/models"
	"ttpham0111/app/web/middleware"
)

type Route struct {
	Name        string
	Methods     []string
	Path        string
	HandlerFunc http.HandlerFunc
}

func GetRoutes(datastore middleware.Datastore, queue chan *models.Order) []Route {
	ordersController := controllers.NewOrdersController(datastore)
	paymentController := controllers.NewPaymentController(datastore, queue)

	return []Route{
		Route{
			"status",
			[]string{http.MethodGet},
			"/status",
			func(w http.ResponseWriter, r *http.Request) {
				status := struct {
					Status string `json:"status"`
				}{"OK"}

				if err := json.NewEncoder(w).Encode(status); err != nil {
					panic(err)
				}
			},
		},

		Route{
			"get-orders",
			[]string{http.MethodGet},
			"/orders",
			ordersController.GetOrders,
		},

		Route{
			"get-order",
			[]string{http.MethodGet},
			"/order/{order-id}",
			ordersController.GetOrder,
		},

		Route{
			"create-order",
			[]string{http.MethodPost},
			"/order",
			ordersController.CreateOrder,
		},

		Route{
			"update-order",
			[]string{http.MethodPut},
			"/order/{order-id}",
			ordersController.UpdateOrder,
		},

		Route{
			"delete-order",
			[]string{http.MethodDelete},
			"/order/{order-id}",
			ordersController.DeleteOrder,
		},

		Route{
			"create-payment",
			[]string{http.MethodPost},
			"/order/{order-id}/pay",
			paymentController.CreatePayment,
		},
	}
}
