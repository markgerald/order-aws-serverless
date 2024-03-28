package response

import "github.com/markgerald/vw-order/model"

type OrdersResponse struct {
	ID      string            `json:"id"`
	Total   float32           `json:"total"`
	UserID  int               `json:"user_id"`
	IsPayed bool              `json:"payed"`
	Items   []model.OrderItem `json:"items"`
}
