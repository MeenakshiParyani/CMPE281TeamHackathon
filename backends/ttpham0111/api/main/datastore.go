package main

import "fmt"

type NoResultFound struct{}

func (err NoResultFound) Error() string {
	return fmt.Sprintf("Not Found")
}

type Datastore interface {
	GetOrders() (error, []Order)
	GetOrder() (error, Order)
}
