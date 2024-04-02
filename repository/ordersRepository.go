package repository

import (
	"github.com/guregu/dynamo"
	"github.com/markgerald/vw-order/model"
	"log"
)

type OrdersRepositoryImpl struct {
	Db *dynamo.DB
}

func NewOrdersRepositoryImpl(Db *dynamo.DB) *OrdersRepositoryImpl {
	return &OrdersRepositoryImpl{
		Db: Db,
	}
}

func (r *OrdersRepositoryImpl) Save(orders model.Order) {
	table := r.Db.Table("orders")
	table.Put(orders).Run()
}

func (r *OrdersRepositoryImpl) Update(orders model.Order) {
	table := r.Db.Table("orders")
	err := table.Put(orders).Run()
	if err != nil {
		return
	}
}

func (r *OrdersRepositoryImpl) Delete(orderId string) {
	table := r.Db.Table("orders")
	table.Delete("id", orderId).Run()
}

func (r *OrdersRepositoryImpl) FindAll() []model.Order {
	var orders []model.Order
	table := r.Db.Table("orders")
	err := table.Scan().All(&orders)
	if err != nil {
		log.Printf("Error fetching all orders: %v", err)
		return nil
	}
	return orders
}

func (r *OrdersRepositoryImpl) FindById(id string) (model.Order, error) {
	var order model.Order
	table := r.Db.Table("orders")
	err := table.Get("ID", id).One(&order)
	if err != nil {
		return order, err
	}
	return order, nil
}
