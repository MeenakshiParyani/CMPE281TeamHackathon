package middleware

import (
	"github.com/satori/go.uuid"
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
	for i, _ := range orders {
		if orders[i].ID == orderID {
			return nil, &orders[i]
		}
	}

	return NoResultFound{}, nil
}

func (d *LocalDatastore) CreateOrder(order *models.Order) (error, *models.Order) {
	order.ID = uuid.NewV4().String()
	orders = append(orders, *order)
	return nil, &orders[len(orders)-1]
}

func (d *LocalDatastore) UpdateOrder(orderID string, newOrder *models.Order) (error, *models.Order) {
	for i, _ := range orders {
		if orders[i].ID == orderID {
			order := orders[i]
			order.Location = newOrder.Location
			order.Items = newOrder.Items
			return nil, &order
		}
	}

	return NoResultFound{}, nil
}

func (d *LocalDatastore) DeleteOrder(orderID string) error {
	lastIndex := len(orders) - 1
	for i, order := range orders {
		if order.ID == orderID {
			orders[i] = orders[lastIndex]
			orders[lastIndex] = models.Order{}
			orders = orders[:lastIndex]
			return nil
		}
	}

	return NoResultFound{}
}
