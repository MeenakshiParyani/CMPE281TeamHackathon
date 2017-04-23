package middleware

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ttpham0111/app/models"
)

type MgoDatastore struct {
	s *mgo.Session
	c *mgo.Collection
}

func NewMgoDatastore(dbUrl string) *MgoDatastore {
	s, err := mgo.Dial(dbUrl)
	if err != nil {
		panic(err)
	}

	c := s.DB("").C("orders")
	return &MgoDatastore{s, c}
}

func (d *MgoDatastore) Close() {
	d.s.Close()
}

func (d *MgoDatastore) GetOrders() (error, []models.Order) {
	var orders []models.Order
	err := d.c.Find(nil).All(&orders)
	if err != nil {
		panic(err)
	}
	return nil, orders
}

func (d *MgoDatastore) GetOrder(orderID string) (error, *models.Order) {
	var order models.Order
	err := d.c.FindId(bson.ObjectIdHex(orderID)).One(&order)
	if err != nil {
		panic(err)
	}
	return nil, &order
}

func (d *MgoDatastore) CreateOrder(order *models.Order) (error, *models.Order) {
	err := d.c.Insert(order)
	if err != nil {
		panic(err)
	}
	return nil, order
}

func (d *MgoDatastore) UpdateOrder(orderID string, order *models.Order) (error, *models.Order) {
	err := d.c.UpdateId(bson.ObjectIdHex(orderID), order)
	if err != nil {
		panic(err)
	}
	return nil, order
}

func (d *MgoDatastore) DeleteOrder(orderID string) error {
	err := d.c.RemoveId(bson.ObjectIdHex(orderID))
	return err
}
