// TODO: Fill in links

package models

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
	ID       string            `json:"id"`
	Location string            `json:"location"`
	Items    []Item            `json:"items"`
	Links    map[string]string `json:"links"`
	Message  string            `json:"message"`
	Status   int               `json:"status"`
}

func (order *Order) SetOrderStatus(status int, uri string) {
	switch status {
	case OrderPlaced:
		order.Status = OrderPlaced
		order.Message = "Order has been placed."
		order.Links = map[string]string{
			"order":   uri + "/order/" + order.ID,
			"payment": uri + "/order/" + order.ID + "/pay",
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
