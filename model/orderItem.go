package model

type OrderItem struct {
	Name   string  `dynamo:"name"`
	Price  float32 `dynamo:"price"`
	Amount int     `dynamo:"amount"`
}
