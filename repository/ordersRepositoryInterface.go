package repository

import "github.com/markgerald/vw-order/model"

type OrdersRepositoryInterface interface {
	Save(orders model.Order)
	Update(orders model.Order)
	Delete(orderId string)
	FindAll() []model.Order
	FindById(id string) (orders *model.Order, err error)
}
