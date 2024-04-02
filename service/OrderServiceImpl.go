package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/markgerald/vw-order/data/request"
	"github.com/markgerald/vw-order/data/response"
	"github.com/markgerald/vw-order/helper"
	"github.com/markgerald/vw-order/model"
	"github.com/markgerald/vw-order/repository"
)

type OrderServiceImpl struct {
	OrdersRepository repository.OrdersRepositoryInterface
	Validate         *validator.Validate
}

func NewOrderServiceImpl(ordersRepository repository.OrdersRepositoryInterface, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrdersRepository: ordersRepository,
		Validate:         validate,
	}
}

func (o OrderServiceImpl) Create(orders request.CreateOrdersRequest) {
	err := o.Validate.Struct(orders)
	helper.ErrorPanic(err)
	orderModel := model.Order{
		Total:   orders.Total,
		UserID:  orders.UserID,
		IsPayed: orders.IsPayed,
		Items:   orders.Items,
	}
	o.OrdersRepository.Save(SumOrder(orderModel))
}

func (o OrderServiceImpl) Update(orders request.UpdateOrdersRequest) response.OrdersResponse {
	orderData, err := o.OrdersRepository.FindById(orders.ID)
	helper.ErrorPanic(err)

	orderResponse := response.OrdersResponse{
		ID:      orderData.ID,
		Total:   orderData.Total,
		UserID:  orderData.UserID,
		IsPayed: orderData.IsPayed,
		Items:   orderData.Items,
	}

	return orderResponse
}

func (o OrderServiceImpl) Delete(orderId string) {
	o.OrdersRepository.Delete(orderId)
}

func (o OrderServiceImpl) FindByID(orderId string) (response.OrdersResponse, error) {
	orderData, err := o.OrdersRepository.FindById(orderId)
	helper.ErrorPanic(err)

	orderResponse := response.OrdersResponse{
		ID:      orderData.ID,
		Total:   orderData.Total,
		UserID:  orderData.UserID,
		IsPayed: orderData.IsPayed,
		Items:   orderData.Items,
	}
	return orderResponse, nil
}

func (o OrderServiceImpl) FindAll() []response.OrdersResponse {
	result := o.OrdersRepository.FindAll()

	var orders []response.OrdersResponse
	for _, order := range result {
		orders = append(orders, response.OrdersResponse{
			ID:      order.ID,
			Total:   order.Total,
			UserID:  order.UserID,
			IsPayed: order.IsPayed,
			Items:   order.Items,
		})
	}
	return orders
}
