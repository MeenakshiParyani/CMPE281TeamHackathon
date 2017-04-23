package middleware

import (
	"math"
	"math/rand"
	"time"
	"ttpham0111/app/models"
)

type Barista struct {
	speed float32
	q     chan *models.Order
	d     Datastore
}

func NewBarista(speed float32, q chan *models.Order, d Datastore) *Barista {
	return &Barista{speed, q, d}
}

func (barista *Barista) Start() {
	for order := range barista.q {
		order.SetOrderStatus(models.OrderPreparing, "")
		go barista.d.UpdateOrder(order.ID.Hex(), order)

		prepTime := time.Duration(120/barista.speed) * time.Second
		accidents := time.Duration(rand.Intn(60))
		time.Sleep(prepTime + accidents)

		order.SetOrderStatus(models.OrderServed, "")
		go barista.d.UpdateOrder(order.ID.Hex(), order)

		customerAFK := time.Duration(rand.Intn(120)) * time.Second
		time.Sleep(customerAFK)

		order.SetOrderStatus(models.OrderCollected, "")
		go barista.d.UpdateOrder(order.ID.Hex(), order)

		barista.gainExperience()
	}
}

func (barista *Barista) gainExperience() {
	newSpeed := barista.speed + rand.Float32()
	if newSpeed < barista.speed {
		newSpeed = math.MaxFloat32
	}

	barista.speed = newSpeed
}
