package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"
	"github.com/markgerald/vw-order/count/service"
	"github.com/markgerald/vw-order/db"
	"github.com/markgerald/vw-order/model"
	"log"
	"os"
)

func GetDb() *dynamo.DB {
	return db.InitDb()
}

func HandleRequest() (int, error) {
	var orders []model.Order
	table := GetDb().Table(os.Getenv("DYNAMODB_TABLE"))
	if err := table.Scan().All(&orders); err != nil {
		log.Fatal(err.Error())
	}
	count := len(orders)

	service.PersistFile(count)
	return count, nil
}
func main() {
	lambda.Start(HandleRequest)
}
