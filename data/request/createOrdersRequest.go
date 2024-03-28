package request

import "github.com/markgerald/vw-order/model"

type CreateOrdersRequest struct {
	Total   float32           `json:"total"`
	UserID  int               `json:"user_id"`
	IsPayed bool              `json:"payed"`
	Items   []model.OrderItem `json:"items"`
}
