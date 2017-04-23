package middleware

import (
	"gopkg.in/mgo.v2/bson"
	"ttpham0111/app/models"
)

type LocalDatastore struct{}

var orders = []models.Order{
	models.Order{
		"a",
		"Sunnyvale",
		[]models.Item{
			models.Item{
				2,
				"Mocha Frappuccino",
				"Whole Milk",
				"Grande",
			},
			models.Item{
				1,
				"Caramel Frappuccino",
				"Whole Milk",
				"Grande",
			},
		},
		map[string]string{},
		"No Message",
		models.OrderPlaced,
	},
	models.Order{
		"b",
		"San Jose",
		[]models.Item{
			models.Item{
				1,
				"Mocha Frappuccino",
				"Non-Fat Milk",
				"Grande",
			},
			models.Item{
				3,
				"Caramel Latte",
				"Whole Milk",
				"Venti",
			},
		},
		map[string]string{},
		"No Message",
		models.OrderPreparing,
	},
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
	order.ID = bson.NewObjectId()
	orders = append(orders, *order)
	return nil, &orders[len(orders)-1]
}

func (d *LocalDatastore) UpdateOrder(orderID string, orderRequest *models.OrderRequest) (error, *models.Order) {
	objectID := bson.ObjectIdHex(orderID)
	for i, _ := range orders {
		if orders[i].ID == objectID {
			order := orders[i]
			order.Location = orderRequest.Location
			order.Items = orderRequest.Items
			return nil, &order
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
