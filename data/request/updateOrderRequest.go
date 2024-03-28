package request

type UpdateOrdersRequest struct {
	ID        string `validate:"required" json:"id"`
	ProductID string `validate:"required" json:"product_id"`
	Quantity  int    `validate:"required" json:"quantity"`
}
