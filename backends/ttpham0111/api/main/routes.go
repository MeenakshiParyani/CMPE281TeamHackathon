package main

import (
	"encoding/json"
	"net/http"
)

type Route struct {
	Name        string
	Methods     []string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
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
		GetOrders,
	},

	Route{
		"get-order",
		[]string{http.MethodGet},
		"/order/{order-id}",
		GetOrder,
	},

	Route{
		"create-order",
		[]string{http.MethodPost},
		"/order",
		CreateOrder,
	},

	Route{
		"update-order",
		[]string{http.MethodPut},
		"/order/{order-id}",
		UpdateOrder,
	},

	Route{
		"delete-order",
		[]string{http.MethodDelete},
		"/order/{order-id}",
		DeleteOrder,
	},

	Route{
		"create-payment",
		[]string{http.MethodPost},
		"/order/{order-id}/pay",
		CreatePayment,
	},
}
