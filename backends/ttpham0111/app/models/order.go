package models

import "gopkg.in/mgo.v2/bson"

const (
	OrderPlaced    = "Placed"
	OrderPaid      = "Paid"
	OrderPreparing = "Preparing"
	OrderServed    = "Served"
	OrderCollected = "Colleged"
)

type Item struct {
	Quantity int    `json:"qty"`
	Name     string `json:"name"`
	Milk     string `json:"milk"`
	Size     string `json:"size"`
}

type Order struct {
	ID       bson.ObjectId     `json:"_id" bson:"_id,omitempty"`
	Location string            `json:"location"`
	Items    []Item            `json:"items"`
	Links    map[string]string `json:"links"`
	Message  string            `json:"message"`
	Status   string            `json:"status"`
}

type OrderRequest struct {
	Location string `json:"location"`
	Items    []Item `json:"items"`
}

func (order *Order) SetOrderStatus(status string, uri string) {
	switch status {
	case OrderPlaced:
		order.Status = OrderPlaced
		order.Message = "Order has been placed."
		order.Links = map[string]string{
			"order":   uri + "/order/" + order.ID.Hex(),
			"payment": uri + "/order/" + order.ID.Hex() + "/pay",
		}

	case OrderPaid:
		order.Status = OrderPaid
		order.Message = "Payment accepted."
		delete(order.Links, "payment")

	case OrderPreparing:
		order.Status = OrderPreparing
		order.Message = "Order preparations in progress."

	case OrderServed:
		order.Status = OrderServed
		order.Message = "Order served, wating for Customer pickup."

	case OrderCollected:
		order.Status = OrderCollected
		order.Message = "Order retrived by Customer."
	}
}
