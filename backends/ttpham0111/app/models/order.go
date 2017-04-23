package models

import "gopkg.in/mgo.v2/bson"

const (
	OrderPlaced = iota
	OrderPaid
	OrderPreparing
	OrderServed
	OrderCollected
)

type Item struct {
	Quantity int    `json:"qty"`
	Name     string `json:"name"`
	Milk     string `json:"milk"`
	Size     string `json:"size"`
}

type Order struct {
	ID       bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	Location string            `json:"location"`
	Items    []Item            `json:"items"`
	Links    map[string]string `json:"links"`
	Message  string            `json:"message"`
	Status   int               `json:"-"`
}

type OrderRequest struct {
	Location string `json:"location"`
	Items    []Item `json:"items"`
}

func (order *Order) SetOrderStatus(status int, uri string) {
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
