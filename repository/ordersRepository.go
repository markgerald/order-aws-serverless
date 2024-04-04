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
	table := r.Db.Table("orders-prod")
	table.Put(orders).Run()
}

func (r *OrdersRepositoryImpl) Update(order model.Order) (error error) {
	table := r.Db.Table("orders-prod")
	err := table.Put(order).Run()
	if err != nil {
		log.Printf("Error updating order: %v", err)
		return err
	}
	return nil
}

func (r *OrdersRepositoryImpl) Delete(orderId string) (error error) {
	table := r.Db.Table("orders-prod")
	err := table.Delete("id", orderId).Run()
	if err != nil {
		log.Printf("Error deleting order: %v", err)
		return err
	}
	return nil
}

func (r *OrdersRepositoryImpl) FindAll() []model.Order {
	var orders []model.Order
	table := r.Db.Table("orders-prod")
	err := table.Scan().All(&orders)
	if err != nil {
		log.Printf("Error fetching all orders: %v", err)
		return nil
	}
	return orders
}

func (r *OrdersRepositoryImpl) FindById(id string) (*model.Order, error) {
	var order model.Order
	table := r.Db.Table("orders-prod")
	err := table.Get("id", id).One(&order)
	if err != nil {
		log.Printf("Error fetching this order: %v", err)
		return nil, err
	}
	return &order, nil
}
