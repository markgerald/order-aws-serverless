package service

import (
	"github.com/google/uuid"
	"github.com/markgerald/vw-order/model"
)

func SumOrder(order model.Order) model.Order {
	order.ID = uuid.New().String()
	for _, item := range order.Items {
		for i := 0; i < item.Amount; i++ {
			order.Total += item.Price
		}
	}

	return order
}
