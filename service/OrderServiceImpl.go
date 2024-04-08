package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/markgerald/vw-order/data/request"
	"github.com/markgerald/vw-order/data/response"
	"github.com/markgerald/vw-order/helper"
	"github.com/markgerald/vw-order/model"
	"github.com/markgerald/vw-order/repository"
	"log"
)

type OrderServiceImpl struct {
	OrdersRepository repository.OrdersRepositoryInterface
	Validate         *validator.Validate
}

func NewOrderServiceImpl(
	ordersRepository repository.OrdersRepositoryInterface,
	validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrdersRepository: ordersRepository,
		Validate:         validate,
	}
}

func (o OrderServiceImpl) Create(orders request.CreateOrdersRequest) (*response.OrdersResponse, error) {
	err := o.Validate.Struct(orders)
	helper.ErrorPanic(err)
	orderModel := model.Order{
		Total:   orders.Total,
		UserID:  orders.UserID,
		IsPayed: orders.IsPayed,
		Items:   orders.Items,
	}
	save, err := o.OrdersRepository.Save(SumOrder(orderModel))
	if err != nil {
		return nil, err
	}

	return &response.OrdersResponse{
		ID:      save.ID,
		Total:   save.Total,
		UserID:  save.UserID,
		IsPayed: save.IsPayed,
		Items:   save.Items,
	}, nil
}

func (o OrderServiceImpl) Update(orders request.UpdateOrdersRequest) (*response.OrdersResponse, error) {
	orderData, err := o.OrdersRepository.FindById(orders.ID)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}

	orderResponse := response.OrdersResponse{
		ID:      orderData.ID,
		Total:   orderData.Total,
		UserID:  orderData.UserID,
		IsPayed: orderData.IsPayed,
		Items:   orderData.Items,
	}

	o.OrdersRepository.Update(SumOrder(*orderData))

	return &orderResponse, nil
}

func (o OrderServiceImpl) Delete(orderId string) (error error) {
	err := o.OrdersRepository.Delete(orderId)
	if err != nil {
		return err
	}
	return nil
}

func (o OrderServiceImpl) FindByID(orderId string) (*response.OrdersResponse, error) {
	orderData, err := o.OrdersRepository.FindById(orderId)
	if err != nil {
		return nil, err
	}

	orderResponse := response.OrdersResponse{
		ID:      orderData.ID,
		Total:   orderData.Total,
		UserID:  orderData.UserID,
		IsPayed: orderData.IsPayed,
		Items:   orderData.Items,
	}
	return &orderResponse, nil
}

func (o OrderServiceImpl) FindAll(limit int, startKey string) ([]response.OrdersResponse, string, error) {
	result, lastKey, err := o.OrdersRepository.FindAll(limit, startKey)
	if err != nil {
		return nil, "", err
	}

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
	return orders, lastKey, nil
}

func (s *OrderServiceImpl) FindByUserId(userId int, limit string, startKey string) ([]response.OrdersResponse, string, error) {
	result, lastKey, err := s.OrdersRepository.FindByUserId(userId, limit, startKey)
	if err != nil {
		return nil, "", err
	}
	var ordersResponse []response.OrdersResponse
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

	return ordersResponse, lastKey, nil
}
