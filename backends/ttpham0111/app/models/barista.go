package models

import (
	"math/rand"
	"time"
)

type Barista struct {
	Queue chan *Order
	Speed float32
}

func (barista Barista) Start() {
	for order := range barista.Queue {
		order.SetOrderStatus(OrderPreparing)

		prepTime := time.Duration(120/barista.Speed) * time.Second
		accidents := time.Duration(rand.Intn(60))
		time.Sleep(prepTime + accidents)

		order.SetOrderStatus(OrderServed)

		customerAFK := time.Duration(rand.Intn(120)) * time.Second
		time.Sleep(customerAFK)

		order.SetOrderStatus(OrderCollected)
	}
}
