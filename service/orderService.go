package service

import (
	"github.com/markgerald/vw-order/data/request"
	"github.com/markgerald/vw-order/data/response"
)

type OrderService interface {
	Create(orders request.CreateOrdersRequest)
	Update(orders request.UpdateOrdersRequest) response.OrdersResponse
	Delete(orderId string)
	FindByID(orderId string) response.OrdersResponse
	FindAll() []response.OrdersResponse
}
