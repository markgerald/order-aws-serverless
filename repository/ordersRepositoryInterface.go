package repository

import "github.com/markgerald/vw-order/model"

type OrdersRepositoryInterface interface {
	Save(orders model.Order)
	Update(orders model.Order)
	Delete(orderId string) (error error)
	FindAll() []model.Order
	FindById(id string) (order *model.Order, err error)
}
