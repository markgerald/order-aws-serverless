package main

import (
	"github.com/guregu/dynamo"
	"github.com/markgerald/vw-order/db"
	"github.com/markgerald/vw-order/model"
	"log"
	"os"
)

func GetDb() *dynamo.DB {
	return db.InitDb()
}

func main() {
	var orders []model.Order
	table := GetDb().Table(os.Getenv("DYNAMODB_TABLE"))
	if err := table.Scan().All(&orders); err != nil {
		log.Fatal(err.Error())
		return
	}
	count := len(orders)

	PersistFile(count)
}
