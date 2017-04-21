package main

import "github.com/satori/go.uuid"

type LocalDatastore struct{}

var orders = []Order{
	Order{
		"a",
		"Sunnyvale",
		[]Item{
			Item{
				2,
				"Mocha Frappuccino",
				"Whole Milk",
				"Grande",
			},
			Item{
				1,
				"Caramel Frappuccino",
				"Whole Milk",
				"Grande",
			},
		},
		map[string]string{},
		"No Message",
		PLACED,
	},
	Order{
		"b",
		"San Jose",
		[]Item{
			Item{
				1,
				"Mocha Frappuccino",
				"Non-Fat Milk",
				"Grande",
			},
			Item{
				3,
				"Caramel Latte",
				"Whole Milk",
				"Venti",
			},
		},
		map[string]string{},
		"No Message",
		PREPARING,
	},
}

func (d *LocalDatastore) GetOrders() (error, []Order) {
	return nil, orders
}

func (d *LocalDatastore) GetOrder(orderID string) (error, *Order) {
	for i, _ := range orders {
		if orders[i].ID == orderID {
			return nil, &orders[i]
		}
	}

	return NoResultFound{}, nil
}

func (d *LocalDatastore) CreateOrder(order *Order) (error, *Order) {
	order.ID = uuid.NewV4().String()
	orders = append(orders, *order)
	return nil, &orders[len(orders)-1]
}

func (d *LocalDatastore) UpdateOrder(orderID string, newOrder *Order) (error, *Order) {
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
			orders[lastIndex] = Order{}
			orders = orders[:lastIndex]
			return nil
		}
	}

	return NoResultFound{}
}
