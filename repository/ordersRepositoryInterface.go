package repository

import "github.com/markgerald/vw-order/model"

type OrdersRepositoryInterface interface {
	Save(orders model.Order)
	Update(orders model.Order) (error error)
	Delete(orderId string) (error error)
	FindAll(limit int, startKey string) ([]model.Order, string, error)
	FindById(id string) (order *model.Order, err error)
}
