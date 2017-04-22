package middleware

import (
	"fmt"
	"ttpham0111/app/models"
)

type NoResultFound struct{}

func (err NoResultFound) Error() string {
	return fmt.Sprintf("Not Found")
}

type Datastore interface {
	GetOrders() (error, []models.Order)
	GetOrder(string) (error, *models.Order)
	CreateOrder(*models.OrderRequest) (error, *models.Order)
	UpdateOrder(string, *models.OrderRequest) (error, *models.Order)
	DeleteOrder(string) error
}
