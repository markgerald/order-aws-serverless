package model

type Order struct {
	ID      string      `dynamo:"id"`
	Total   float32     `dynamo:"total"`
	UserID  int         `dynamo:"user_id"`
	IsPayed bool        `dynamo:"payed"`
	Items   []OrderItem `dynamo:"items"`
}
