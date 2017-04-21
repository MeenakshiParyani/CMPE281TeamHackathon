// TODO: Fill in links

package main

const (
	PLACED = iota
	PAID
	PREPARING
	SERVED
	COLLECTED
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
	status   int
}

func (order *Order) SetOrderStatus(status int) {
	switch status {
	case PLACED:
		order.status = PLACED
		order.Message = "Order has been placed."
		order.Links = map[string]string{
			"orders":  "",
			"payment": "",
		}

	case PAID:
		order.status = PAID
		order.Message = "Payment accepted."
		delete(order.Links, "payment")

	case PREPARING:
		order.status = PREPARING
		order.Message = "Order preparations in progress."

	case SERVED:
		order.status = SERVED
		order.Message = "Order served, wating for Customer pickup."

	case COLLECTED:
		order.status = COLLECTED
		order.Message = "Order retrived by Customer."
	}
}
