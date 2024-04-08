package service

import (
	"github.com/markgerald/vw-order/data/request"
	"github.com/markgerald/vw-order/data/response"
)

type OrderService interface {
	Create(orders request.CreateOrdersRequest) (*response.OrdersResponse, error)
	Update(orders request.UpdateOrdersRequest) (*response.OrdersResponse, error)
	Delete(orderId string) (error error)
	FindByID(orderId string) (*response.OrdersResponse, error)
	FindAll(limit int, startKey string) ([]response.OrdersResponse, string, error)
	FindByUserId(userId int, limit string, startKey string) ([]response.OrdersResponse, string, error)
}
