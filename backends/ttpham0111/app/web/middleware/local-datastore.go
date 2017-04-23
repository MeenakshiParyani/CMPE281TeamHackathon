package middleware

import (
	"gopkg.in/mgo.v2/bson"
	"ttpham0111/app/models"
)

type LocalDatastore struct{}

var orders []models.Order = make([]models.Order, 0)

func (d *LocalDatastore) Close() {
	return
}

func (d *LocalDatastore) GetOrders() (error, []models.Order) {
	return nil, orders
}

func (d *LocalDatastore) GetOrder(orderID string) (error, *models.Order) {
	objectID := bson.ObjectIdHex(orderID)
	for i, _ := range orders {
		if orders[i].ID == objectID {
			return nil, &orders[i]
		}
	}

	return NoResultFound{}, nil
}

func (d *LocalDatastore) CreateOrder(order *models.Order) (error, *models.Order) {
	orders = append(orders, *order)
	return nil, &orders[len(orders)-1]
}

func (d *LocalDatastore) UpdateOrder(orderID string, order *models.Order) (error, *models.Order) {
	objectID := bson.ObjectIdHex(orderID)
	for i, _ := range orders {
		if orders[i].ID == objectID {
			orders[i] = *order
			return nil, order
		}
	}

	return NoResultFound{}, nil
}

func (d *LocalDatastore) DeleteOrder(orderID string) error {
	objectID := bson.ObjectIdHex(orderID)
	lastIndex := len(orders) - 1
	for i, order := range orders {
		if order.ID == objectID {
			orders[i] = orders[lastIndex]
			orders[lastIndex] = models.Order{}
			orders = orders[:lastIndex]
			return nil
		}
	}

	return NoResultFound{}
}
